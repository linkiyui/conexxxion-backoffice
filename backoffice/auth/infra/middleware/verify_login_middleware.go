package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/auth/infra/auth_token"
)

func VerifyLoginToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := auth_token.ValidateLoginToken(strings.TrimPrefix(token, "Bearer "))
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	c.Set("user_id", claims.UserID)
	c.Set("role", claims.Role)
	c.Next()
}
