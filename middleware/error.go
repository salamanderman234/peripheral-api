package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/entity"
	utility "github.com/salamanderman234/peripheral-api/utility"
)

func Error(err error, ctx echo.Context) {
	report, ok := err.(*echo.HTTPError)

	if ok {
		report.Message = fmt.Sprintf("%v", report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	utility.NewLogEntry(ctx).Error(fmt.Sprintf("%d - %s", report.Code, report.Message))
	ctx.JSON(report.Code, entity.BaseResponse{
		Code:   report.Code,
		Status: report.Message.(string),
	})
}
