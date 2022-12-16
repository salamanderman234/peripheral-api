package route

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/config"
	"github.com/salamanderman234/peripheral-api/domain"
)

type switchRoute struct {
	router  *echo.Echo
	con     domain.SwitchController
	baseUrl string
}

func NewSwitchRoute(router *echo.Echo, con domain.SwitchController) domain.Router {
	return &switchRoute{
		router:  router,
		con:     con,
		baseUrl: config.CreateAppPath("/switches"),
	}
}

func (s *switchRoute) CreateNewURL(str string) string {
	return fmt.Sprintf("%s%s", s.baseUrl, str)
}

func (s *switchRoute) RegisterRoutes() {
	s.router.GET(s.CreateNewURL(""), s.con.GetAllSwitch)
	s.router.GET(s.CreateNewURL("/:slug"), s.con.GetOneSwitch)
	s.router.POST(s.CreateNewURL(""), s.con.CreateNewSwitch)
	s.router.PATCH(s.CreateNewURL("/:slug"), s.con.UpdateOneSwitch)
	s.router.DELETE(s.CreateNewURL("/:slug"), s.con.DropSwitch)
}
