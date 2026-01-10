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

	err := h.recordService.Update(c, &req)
	if err != nil {
		response.Fail(c, response.ServerError)
		return
	}

	response.Success(c, nil)
}
