package model

import (
	"checkme/internal/dto"
	"time"

	"gorm.io/gorm"
)

type Record struct {
	ID          uint           `gorm:"primarykey"`
	Device      string         `json:"device" gorm:"varchar(255)"`
	Application string         `json:"application" gorm:"varchar(255)"`
	UpdatedTime time.Time      `json:"updated_time"`
	StartTime   time.Time      `json:"start_time"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (Record) TableName() string {
	return "records"
}

func (r *Record) ToDeviceRecord() dto.DeviceRecord {
	return dto.DeviceRecord{
		Device:      r.Device,
		Application: r.Application,
		StartTime:   r.StartTime,
		UpdateTime:  r.UpdatedTime,
	}
}
