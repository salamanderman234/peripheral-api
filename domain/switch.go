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
	FindAllSwitchWithFilter(ctx context.Context, switchType string, switchManufacturer string, acforce float64) ([]model.Switch, error)
}

type SwitchService interface {
	GetSwitch(ctx context.Context, filter entity.Switch) ([]byte, error)
	CreateSwitch(ctx context.Context, switchs []entity.Switch) error
}

type SwitchController interface {
	GetAllSwitch(ctx echo.Context) error
	CreateNewSwitch(ctx echo.Context) error
}
