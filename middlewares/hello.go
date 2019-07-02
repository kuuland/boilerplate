package middlewares

import (
	"github.com/gin-gonic/gin"
)

// HelloMiddleware
func HelloMiddleware(c *gin.Context) {
	c.Next()
}
