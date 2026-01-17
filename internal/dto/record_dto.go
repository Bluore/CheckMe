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
	Device     string         `json:"device"`
	StartTime  time.Time      `json:"start_time"`
	UpdateTime time.Time      `json:"update_time"`
	Data       datatypes.JSON `json:"data"`
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

type CreateNotifyRequest struct {
	Type string `json:"type" binding:"required"`
	Msg  string `json:"msg" binding:"max=20"`
}
