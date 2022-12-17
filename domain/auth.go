package domain

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/entity"
	model "github.com/salamanderman234/peripheral-api/models"
)

type AuthRepository interface {
	CreateNewUser(ctx context.Context, user model.User) (model.User, error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
}

type AuthService interface {
	Authenticate(ctx context.Context, cred entity.Credentials) (string, error)
	Register(ctx context.Context, user entity.User) (entity.User, bool, error)
}

type AuthController interface {
	Login(ctx echo.Context) error
	Register(ctx echo.Context) error
}
