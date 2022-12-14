package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/entity"
	utility "github.com/salamanderman234/peripheral-api/utility"
)

func Error(err error, ctx echo.Context) {
	var errorMessage string
	report, ok := err.(*echo.HTTPError)

	if ok {
		report.Message = fmt.Sprintf("%v", report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if report.Code == 404 {
		errorMessage = "The resource you are trying to reach cannot be found on this server"
	} else {
		errorMessage = "Something went wrong"
	}

	utility.NewLogEntry(ctx).Error(fmt.Sprintf("%d - %s", report.Code, report.Message))
	ctx.JSON(report.Code, entity.BaseResponse{
		Code:   report.Code,
		Status: report.Message.(string),
		Errors: errorMessage,
	})
}
