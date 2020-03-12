package middleware

import (
	"github.com/demos/api_server/handler"
	"github.com/demos/api_server/pkg/errno"
	"github.com/demos/api_server/pkg/token"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware is for token vlaid
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
