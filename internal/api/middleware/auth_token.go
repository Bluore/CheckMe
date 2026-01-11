package middleware

import (
	"checkme/config"
	"checkme/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AuthToken(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		type tokenReq struct {
			Token string `json:"token" binding:"required"`
		}
		var token tokenReq
		if err := c.ShouldBindBodyWith(&token, binding.JSON); err != nil {
			response.Fail(c, response.ParamMissing)
			c.Abort()
			return
		}

		if token.Token != cfg.Auth.Token {
			response.Error(c, 403, "身份验证失败")
			c.Abort()
			return
		}

		c.Next()
	}
}
