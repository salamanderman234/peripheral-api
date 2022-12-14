package domain

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/entity"
	model "github.com/salamanderman234/peripheral-api/models"
)

type SwitchRepository interface {
	InsertSwitch(ctx context.Context, newSwitch model.Switch) (model.Switch, error)
	BatchInsertSwitchs(ctx context.Context, switchs []model.Switch) ([]model.Switch, error)
	UpdateSwitch(ctx context.Context, updateField model.Switch, condition model.Switch) error
	DeleteSwitch(ctx context.Context, condition model.Switch) error
	FindAllSwitchWithFilter(ctx context.Context, filter model.Switch) ([]model.Switch, error)
}

type SwitchService interface {
	GetSwitch(ctx echo.Context, filter entity.SwitchFilter) (*entity.BaseResponse, error)
}

type SwitchRouter interface {
	RegisterRoutes()
	CreateNewURL(str string) string
}
