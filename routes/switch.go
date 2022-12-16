package route

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/domain"
)

type switchRoute struct {
	group *echo.Group
	con   domain.SwitchController
	path  string
}

func NewSwitchRoute(group *echo.Group, con domain.SwitchController) domain.Router {
	return &switchRoute{
		group: group,
		con:   con,
		path:  "/switches",
	}
}

func (s *switchRoute) RegisterRoutes() {
	// api/v1/switches
	group := s.group.Group(s.path)
	group.GET("", s.con.GetAllSwitch)
	group.GET("/:switch_id", s.con.GetOneSwitch)
	group.POST("", s.con.CreateNewSwitch)
	group.PATCH("/:switch_id", s.con.UpdateOneSwitch)
	group.DELETE("/:switch_id", s.con.DropSwitch)
}
