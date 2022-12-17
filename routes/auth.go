package route

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/domain"
)

type authRoute struct {
	group *echo.Group
	con   domain.AuthController
	path  string
}

func NewAuthRoute(group *echo.Group, controller domain.AuthController) domain.Router {
	return &authRoute{
		group: group,
		con:   controller,
		path:  "/session",
	}
}

func (a *authRoute) RegisterRoutes() {
	group := a.group.Group(a.path)
	group.POST("", a.con.Login)
}
