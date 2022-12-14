package service

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/domain"
	"github.com/salamanderman234/peripheral-api/entity"
)

type switchService struct {
	switchRepo domain.SwitchRepository
}

func NewSwitchRepository(repo domain.SwitchRepository) domain.SwitchService {
	return &switchService{
		switchRepo: repo,
	}
}

func (s *switchService) GetSwitch(ctx echo.Context, filter entity.SwitchFilter) ([]byte, error) {
	switchs, err := s.switchRepo.FindAllSwitchWithFilter(ctx.Request().Context(), filter.Type, filter.Manufacturer, filter.ActuationForce)
	if err != nil {
		return nil, err
	}

	switchsParshed, _ := json.Marshal(switchs)
	return switchsParshed, nil
}
