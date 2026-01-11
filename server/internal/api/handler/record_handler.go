package handler

import (
	"checkme/internal/dto"
	"checkme/internal/service"
	"checkme/pkg/response"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	recordService service.RecordService
}

func NewHandler(recoderService service.RecordService) *Handler {
	return &Handler{recordService: recoderService}
}

// UploadRecord 创建记录
func (h *Handler) UploadRecord(c *gin.Context) {
	var req dto.UploadRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(req)
		response.Fail(c, response.ParamMissing)
		return
	}

	if req.Time == nil {
		now := time.Now()
		req.Time = &now
	}

	// 仅支持phone、computer
	if req.Device != "phone" && req.Device != "computer" {
		response.Error(c, 400, "不支持的设备")
		return
	}

	err := h.recordService.Update(c, &req)
	if err != nil {
		response.Fail(c, response.ServerError)
		return
	}

	response.Success(c, nil)
}

// GetLastRecord 获取最近在线情况
func (h *Handler) GetLastRecord(c *gin.Context) {
	devices, err := h.recordService.GetLastRecord(c)
	if err != nil {
		response.Fail(c, response.ServerError)
		return
	}

	response.Success(c, devices)
}
