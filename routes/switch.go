package route

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/controller"
	"github.com/salamanderman234/peripheral-api/domain"
)

type switchRoute struct {
	router  *echo.Echo
	con     *controller.SwitchController
	baseUrl string
}

func NewSwitchRoute(router *echo.Echo, con *controller.SwitchController) domain.SwitchRouter {
	return &switchRoute{
		router:  router,
		con:     con,
		baseUrl: "/api/v1/switch",
	}
}

func (s *switchRoute) CreateNewURL(str string) string {
	return fmt.Sprintf("%s%s", s.baseUrl, str)
}

func (s *switchRoute) RegisterRoutes() {
	s.router.GET(s.CreateNewURL(""), s.con.GetAllSwitch)
}
