package middleware

import (
	"github.com/labstack/echo/v4"
	utility "github.com/salamanderman234/peripheral-api/utility"
)

func Log(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		utility.NewLogEntry(ctx).Info("Incoming Request")
		return next(ctx)
	}
}
