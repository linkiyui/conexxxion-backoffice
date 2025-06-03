package api_server

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	clog "gitlab.com/conexxxion/conexxxion-backoffice/logger"
	"gitlab.com/conexxxion/conexxxion-backoffice/utils"
)

func LoggerMiddleware(c *gin.Context) {
	c.Set("request_id", utils.GenerateULID())
	c.Set("client_ip", c.ClientIP())
	c.Set("path", c.Request.URL.Path)
	c.Set("method", c.Request.Method)
	t := time.Now()
	c.Next()
	d := time.Since(t)
	statusCode := fmt.Sprint(c.Writer.Status())
	clog.InfoCtx(c, "request processed", map[string]any{
		"duration": d.String(),
		"status":   statusCode,
	})
	requestsDuration.WithLabelValues(statusCode).Observe(float64(d.Nanoseconds()) / 1_000_000)
}
