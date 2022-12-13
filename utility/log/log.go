package utility

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func NewLogEntry(ctx echo.Context) *logrus.Entry {
	if ctx == nil {
		return logrus.WithFields(logrus.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}
	return logrus.WithFields(logrus.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": ctx.Request().Method,
		"uri":    ctx.Request().URL.String(),
		"ip":     ctx.Request().RemoteAddr,
	})
}
