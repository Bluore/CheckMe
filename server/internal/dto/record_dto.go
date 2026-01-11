package dto

import (
	"time"
)

type UploadRecordRequest struct {
	Device      string     `json:"device" binding:"required"`
	Application string     `json:"application" binding:"required"`
	Time        *time.Time `json:"time"`
}

type DeviceRecord struct {
	Device      string    `json:"device"`
	Application string    `json:"application"`
	StartTime   time.Time `json:"start_time"`
	UpdateTime  time.Time `json:"update_time"`
}

type GetLastRecordResponse struct {
	DeviceList []DeviceRecord `json:"device_list"`
}
