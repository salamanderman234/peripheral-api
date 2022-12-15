package domain

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/entity"
	model "github.com/salamanderman234/peripheral-api/models"
)

type SwitchRepository interface {
	InsertSwitch(ctx context.Context, newSwitch model.Switch) error
	BatchInsertSwitchs(ctx context.Context, switchs []model.Switch) ([]interface{}, error)
	UpdateSwitch(ctx context.Context, updateField model.Switch, filter model.Switch) (int64, error)
	DeleteSwitch(ctx context.Context, condition model.Switch) error
	FindAllSwitchWithFilter(ctx context.Context, switchType string, switchManufacturer string, acforce float64, slug string) ([]model.Switch, error)
}

type SwitchService interface {
	GetSwitch(ctx context.Context, filter entity.Switch) ([]byte, error)
	CreateSwitch(ctx context.Context, switchs []entity.Switch) ([]interface{}, error)
	CreateOneSwitch(ctx context.Context, switchEntity entity.Switch) error
	UpdateSwitch(ctx context.Context, updateField entity.Switch, filter entity.Switch) (int64, error)
}

type SwitchController interface {
	GetAllSwitch(ctx echo.Context) error
	GetOneSwitch(ctx echo.Context) error
	CreateNewSwitch(ctx echo.Context) error
	UpdateOneSwitch(ctx echo.Context) error
}
