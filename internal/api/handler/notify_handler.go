package handler

import (
	"checkme/internal/dto"
	"checkme/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) CreateNotify(c *gin.Context) {
	var req dto.CreateNotifyRequest
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		response.Fail(c, response.ParamMissing)
		return
	}

	err := h.notifyService.CreateNotify(c, &req)
	if err != nil {
		response.Error(c, 500, "服务器错误")
		return
	}

	response.Success(c, nil)
}
