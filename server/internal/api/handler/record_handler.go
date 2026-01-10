package handler

import (
	"checkme/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	recordService service.RecordService
}

func NewHandler(recoderService service.RecordService) *Handler {
	return &Handler{recordService: recoderService}
}

// todo 创建记录
func (h *Handler) UploadRecord(c *gin.Context) {

}
