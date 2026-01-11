package middleware

import "github.com/gin-gonic/gin"

// CORS 跨域中间件
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Request-Method", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")

		c.Next()
	}
}
