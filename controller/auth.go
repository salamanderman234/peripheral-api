package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/domain"
)

type authController struct {
	service domain.AuthService
}

func NewAuthController(service domain.AuthService) domain.AuthController {
	return &authController{
		service: service,
	}
}

func (a *authController) Login(ctx echo.Context) error {
	return nil
}
func (a *authController) Register(ctx echo.Context) error {
	return nil
}
