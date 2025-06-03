package middlewares

import (
	"github.com/gin-gonic/gin"
)

func VerifyAdmin(c *gin.Context) {

	r, exists := c.Get("role")
	if !exists {
		c.AbortWithStatus(400)
		return
	}

	if r != "admin" && r != "super_admin" {
		c.AbortWithStatus(403)
		return
	}

	c.Next()
}
