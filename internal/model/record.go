package model

import (
	"checkme/internal/dto"
	"checkme/pkg/judge"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Record struct {
	ID          uint           `gorm:"primarykey"`
	Device      string         `json:"device" gorm:"varchar(255)"`
	Application string         `json:"application" gorm:"varchar(255)"`
	UpdatedTime time.Time      `json:"updated_time"`
	StartTime   time.Time      `json:"start_time"`
	Ip          string         `json:"ip"`
	Data        datatypes.JSON `gorm:"type:json"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (r *Record) TableName() string {
	return "records"
}

func (r *Record) ToDeviceRecord() dto.DeviceRecord {
	res := dto.DeviceRecord{
		Device:     r.Device,
		StartTime:  r.StartTime,
		UpdateTime: r.UpdatedTime,
		Data:       r.Data,
	}
	if judge.IsJSONNull(res.Data) {
		res.Data = datatypes.JSON(`{}`)
	}
	return res
}
