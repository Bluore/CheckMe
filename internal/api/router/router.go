package router

import (
	"checkme/config"
	"checkme/internal/api/handler"
	"checkme/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, h *handler.Handler, cfg *config.Config) {
	// 全局中间件
	r.Use(middleware.CORS())

	// 设置路由
	v1 := r.Group("/api/v1")
	{
		record := v1.Group("/admin")
		record.Use(middleware.AuthToken(cfg))
		{
			record.POST("/record", h.UploadRecord)
		}
		recordGuest := v1.Group("/guest")
		{
			recordGuest.GET("/record", h.GetLastRecord)
			recordGuest.GET("/history", h.GetHistoryRecord)
		}
	}
}
