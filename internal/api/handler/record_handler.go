package handler

import (
	"checkme/internal/dto"
	"checkme/internal/service"
	"checkme/pkg/judge"
	"checkme/pkg/response"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/datatypes"
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
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
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

	// 解决null问题
	if judge.IsJSONNull(req.Data) {
		req.Data = datatypes.JSON(`{}`)
	}

	err := h.recordService.Update(c, &req, c.ClientIP())
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

// GetHistoryRecord 获取在线历史记录
func (h *Handler) GetHistoryRecord(c *gin.Context) {
	res, err := h.recordService.GetHistoryRecord(c)
	if err != nil {
		response.Fail(c, response.ServerError)
		return
	}

	response.Success(c, res)
}
