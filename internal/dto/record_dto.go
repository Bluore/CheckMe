package dto

import (
	"time"

	"gorm.io/datatypes"
)

type UploadRecordRequest struct {
	Device      string         `json:"device" binding:"required"`
	Application string         `json:"application" binding:"required"`
	Time        *time.Time     `json:"time"`
	Data        datatypes.JSON `json:"data"`
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

type DeviceRecordList struct {
	DeviceName string         `json:"device_name"`
	Record     []DeviceRecord `json:"record"`
}

type GetHistoryRecordResponse struct {
	List []DeviceRecordList `json:"list"`
}
