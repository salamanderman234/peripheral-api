package service

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/domain"
	"github.com/salamanderman234/peripheral-api/entity"
	model "github.com/salamanderman234/peripheral-api/models"
	utility "github.com/salamanderman234/peripheral-api/utility/log"
)

type switchService struct {
	switchRepo domain.SwitchRepository
}

func NewSwitchRepository(repo domain.SwitchRepository) domain.SwitchService {
	return &switchService{
		switchRepo: repo,
	}
}

func (s *switchService) GetSwitch(ctx echo.Context, filter entity.SwitchFilter) (*entity.BaseResponse, error) {
	filterModel := model.Switch{
		Type: filter.Type,
	}
	switchs, err := s.switchRepo.FindAllSwitchWithFilter(ctx.Request().Context(), filterModel)
	if err != nil {
		utility.NewLogEntry(ctx).Error(err)
		return &entity.BaseResponse{}, err
	}

	return &entity.BaseResponse{
		Message: "success",
		Data:    switchs,
	}, nil
}
