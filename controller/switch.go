package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/domain"
	"github.com/salamanderman234/peripheral-api/entity"
	utility "github.com/salamanderman234/peripheral-api/utility/log"
)

type SwitchController struct {
	service domain.SwitchService
}

func NewSwitchController(service domain.SwitchService) *SwitchController {
	return &SwitchController{
		service: service,
	}
}

func (s *SwitchController) GetAllSwitch(ctx echo.Context) error {
	result, err := s.service.GetSwitch(ctx, entity.SwitchFilter{})
	if err != nil {
		utility.NewLogEntry(ctx).Error("500 - internal server error")
		return ctx.JSON(http.StatusInternalServerError, entity.BaseResponse{
			Message: "failed",
			Error:   "internal server error",
		})
	}

	return ctx.JSON(http.StatusOK, result)
}
