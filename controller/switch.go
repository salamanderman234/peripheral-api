package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/domain"
	"github.com/salamanderman234/peripheral-api/entity"
	utility "github.com/salamanderman234/peripheral-api/utility"
)

type switchController struct {
	service domain.SwitchService
}

func NewSwitchController(service domain.SwitchService) *switchController {
	return &switchController{
		service: service,
	}
}

func (s *switchController) GetAllSwitch(ctx echo.Context) error {
	// init
	var switchFilter entity.SwitchFilter
	var switchs []entity.Switch
	var response entity.BaseResponse

	// get query for filter
	switchType := ctx.QueryParam("type")
	if switchType != "" {
		switchFilter.Type = switchType
	}
	manufacturer := ctx.QueryParam("manufacturer")
	if manufacturer != "" {
		switchFilter.Type = switchType
	}
	actuationForce, err := strconv.ParseFloat(ctx.QueryParam("actuation_force"), 64)
	if err == nil {
		switchFilter.ActuationForce = actuationForce
	}

	// calling service
	result, err := s.service.GetSwitch(ctx, switchFilter)
	if err != nil {
		utility.NewLogEntry(ctx).Error(err)
		response.Status = "internal server error"
		response.Code = http.StatusInternalServerError
	} else {
		// convert result to array of entitiy.switch
		json.Unmarshal(result, &switchs)

		if len(switchs) == 0 {
			utility.NewLogEntry(ctx).Error("404 - Not Found")
			response.Status = "Not Found"
			response.Code = http.StatusNotFound
			response.Errors = "No matching data for the given query"
		} else {
			utility.NewLogEntry(ctx).Info("200 - Success")
			response.Status = "Ok"
			response.Code = http.StatusOK
			response.Data = switchs
		}
	}
	// sending response
	return ctx.JSON(response.Code, response)
}
