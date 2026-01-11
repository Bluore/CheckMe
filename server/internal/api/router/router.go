package router

import (
	"checkme/config"
	"checkme/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, h *handler.Handler, cfg *config.Config) {
	// todo 全局中间件

	// 设置路由
	v1 := r.Group("/api/v1")
	{
		record := v1.Group("/record")
		{
			record.POST("", h.UploadRecord)
			record.GET("", h.GetLastRecord)
		}
	}
}
