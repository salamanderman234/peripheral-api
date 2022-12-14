package service

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/domain"
	"github.com/salamanderman234/peripheral-api/entity"
	model "github.com/salamanderman234/peripheral-api/models"
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
	filterModel := model.Switch{
		Type: filter.Type,
	}
	switchs, err := s.switchRepo.FindAllSwitchWithFilter(ctx.Request().Context(), filterModel)
	if err != nil {
		return nil, err
	}

	switchsParshed, _ := json.Marshal(switchs)
	return switchsParshed, nil
}
