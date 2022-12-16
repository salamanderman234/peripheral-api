package domain

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/entity"
	model "github.com/salamanderman234/peripheral-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SwitchRepository interface {
	InsertSwitch(ctx context.Context, newSwitch model.Switch) error
	BatchInsertSwitches(ctx context.Context, switchs []model.Switch) ([]interface{}, error)
	UpdateSwitch(ctx context.Context, updateField model.Switch, filter primitive.M) (int64, error)
	DeleteSwitch(ctx context.Context, condition string) (int64, error)
	FindAllSwitchWithFilter(ctx context.Context, filter model.Switch, sort string) ([]model.Switch, error)
}

type SwitchService interface {
	GetSwitch(ctx context.Context, filter entity.Switch, sort string) ([]entity.Switch, error)
	CreateSwitch(ctx context.Context, switchs []entity.Switch) ([]interface{}, error)
	CreateOneSwitch(ctx context.Context, switchEntity entity.Switch) error
	UpdateSwitch(ctx context.Context, updateField entity.Switch, filter entity.Switch) (int64, error)
	DeleteSwitch(ctx context.Context, filter string) (int64, error)
}

type SwitchController interface {
	GetAllSwitch(ctx echo.Context) error
	GetOneSwitch(ctx echo.Context) error
	CreateNewSwitch(ctx echo.Context) error
	UpdateOneSwitch(ctx echo.Context) error
	DropSwitch(ctx echo.Context) error
}
