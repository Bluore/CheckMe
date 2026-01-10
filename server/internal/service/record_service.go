package service

import (
	"checkme/config"
	"checkme/internal/dto"
	"checkme/internal/model"
	"checkme/internal/repository"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type RecordService interface {
	Update(ctx context.Context, req *dto.UploadRecordRequest) error
}

type recordService struct {
	recordRepo repository.RecordRepository
	cfg        *config.Config
}

func NewRecoderService(recoderRepo repository.RecordRepository, cfg *config.Config) RecordService {
	return &recordService{
		recordRepo: recoderRepo,
		cfg:        cfg,
	}
}

func (rs recordService) Update(ctx context.Context, req *dto.UploadRecordRequest) error {
	record, err := rs.recordRepo.GetLastByDevice(ctx, req.Device)
	if err != nil {
		return errors.New("数据库查询出错")
	}

	// 更新记录
	if record.Application == req.Application {
		// 当仍在上个应用时更新时长
		record.UpdatedTime = *req.Time
		err := rs.recordRepo.Update(ctx, record)
		if err != nil {
			return err
		}
	} else {
		// 切换应用时添加记录
		rec := model.Record{
			Device:      req.Device,
			Application: req.Application,
			UpdatedTime: *req.Time,
			StartTime:   record.UpdatedTime,
			DeletedAt:   gorm.DeletedAt{},
		}
		if rec.UpdatedTime.Sub(rec.StartTime).Minutes() > 5.0 {
			rec.StartTime = rec.UpdatedTime.Add(-5 * time.Minute)
		}
		err := rs.recordRepo.Create(ctx, &rec)
		if err != nil {
			return err
		}
	}
	return nil
}
