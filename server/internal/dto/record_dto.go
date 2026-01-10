package dto

import "time"

type UploadRecordRequest struct {
	Device      string     `json:"device" binding:"required"`
	Application string     `json:"application" binding:"required"`
	Time        *time.Time `json:"time"`
}
