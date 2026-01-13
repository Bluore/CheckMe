package service

import (
	"checkme/config"
	"checkme/internal/dto"
	"checkme/internal/model"
	"checkme/internal/repository"
	"checkme/pkg/change"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type RecordService interface {
	Update(ctx context.Context, req *dto.UploadRecordRequest, ip string) error
	GetLastRecord(ctx context.Context) (*dto.GetLastRecordResponse, error)
	GetHistoryRecord(ctx context.Context) (*dto.GetHistoryRecordResponse, error)
}

type recordService struct {
	recordRepo repository.RecordRepository
	cfg        *config.Config
	cov        *map[string]map[string]interface{}
}

func NewRecoderService(recoderRepo repository.RecordRepository, cfg *config.Config, cov *map[string]map[string]interface{}) RecordService {
	return &recordService{
		recordRepo: recoderRepo,
		cfg:        cfg,
		cov:        cov,
	}
}

func (rs recordService) Update(ctx context.Context, req *dto.UploadRecordRequest, ip string) error {
	record, err := rs.recordRepo.GetLastByDevice(ctx, req.Device)
	if err != nil {
		return errors.New("数据库查询出错")
	}

	data, err := change.ChangeData(ctx, rs.cov, req, ip)
	if err != nil {
		return err
	}

	// 更新记录
	if record != nil && record.Application == req.Application &&
		req.Time.Sub(record.UpdatedTime).Minutes() < 5.0 {
		// 当仍在上个应用时更新时长
		record.UpdatedTime = *req.Time
		if record.UpdatedTime.Sub(record.StartTime).Minutes() < 0 {
			return errors.New("时间错误，请检查时区")
		}

		record.Ip = ip
		record.Data = data

		err := rs.recordRepo.Update(ctx, record)
		if err != nil {
			return err
		}
	} else {
		// 切换应用时添加记录
		startTime := req.Time.Add(-5 * time.Minute)
		if record != nil {
			if req.Time.Sub(record.UpdatedTime).Minutes() <= 5.0 {
				startTime = record.UpdatedTime
			} else if req.Time.Sub(record.StartTime).Minutes() < 0 {
				return errors.New("时间错误，请检查时区")
			}
		}

		rec := model.Record{
			Device:      req.Device,
			Application: req.Application,
			UpdatedTime: *req.Time,
			StartTime:   startTime,
			Ip:          ip,
			Data:        data,
			DeletedAt:   gorm.DeletedAt{},
		}

		if rec.UpdatedTime.Sub(rec.StartTime).Minutes() < 0 {
			return errors.New("时间错误，请检查时区")
		}

		err := rs.recordRepo.Create(ctx, &rec)
		if err != nil {
			return err
		}
	}
	return nil
}

func (rs recordService) GetLastRecord(ctx context.Context) (*dto.GetLastRecordResponse, error) {
	var res dto.GetLastRecordResponse
	// 查询phone和computer设备
	for _, v := range []string{"phone", "computer"} {
		device, err := rs.recordRepo.GetLastByDevice(ctx, v)
		if err != nil {
			return nil, err
		}
		if device == nil {
			res.DeviceList = append(res.DeviceList,
				dto.DeviceRecord{
					Device:     v,
					StartTime:  time.Time{},
					UpdateTime: time.Time{},
				},
			)
		} else {
			res.DeviceList = append(res.DeviceList,
				device.ToDeviceRecord(),
			)
		}
	}

	return &res, nil
}

func (rs recordService) GetHistoryRecord(ctx context.Context) (*dto.GetHistoryRecordResponse, error) {
	var res dto.GetHistoryRecordResponse
	for _, v := range []string{"phone", "computer"} {
		records, err := rs.recordRepo.GetAllByDeviceAfterDate(ctx, v, time.Now().Add(-2*time.Hour))
		if err != nil {
			return nil, err
		}

		//var devicLis []dto.DeviceRecord
		devicLis := make([]dto.DeviceRecord, 0)
		for _, rec := range records {
			devicLis = append(devicLis, rec.ToDeviceRecord())
		}

		res.List = append(res.List,
			dto.DeviceRecordList{
				DeviceName: v,
				Record:     devicLis,
			},
		)
	}

	return &res, nil
}
