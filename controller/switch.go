package controller

import (
	"encoding/json"
	"net/http"

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
	var switchFilter entity.Switch
	var switchs []entity.Switch
	var response entity.BaseResponse

	// get filter from query, body or path params
	ctx.Bind(&switchFilter)

	// calling service
	result, err := s.service.GetSwitch(ctx.Request().Context(), switchFilter)
	if err != nil {
		utility.NewLogEntry(ctx).Error(err)
		response.Status = "Internal Server Error"
		response.Code = http.StatusInternalServerError
		response.Errors = "Something Went Wrong"
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

func (s *switchController) CreateNewSwitch(ctx echo.Context) error {
	// init
	var switchsBody []entity.Switch
	var response entity.BaseResponse

	// binding
	if err := ctx.Bind(&switchsBody); err != nil {
		utility.NewLogEntry(ctx).Error("400 - Bad Request")
		response.Status = "Bad Request"
		response.Code = http.StatusBadRequest
		response.Errors = "Data body does not match specifications"

	} else {

		// calling service
		err := s.service.CreateSwitch(ctx.Request().Context(), switchsBody)
		// error while parsing body
		if err != nil {
			utility.NewLogEntry(ctx).Error("500 - Internal Server Error")
			response.Status = "internal Server Error"
			response.Code = http.StatusBadRequest
			response.Errors = "Something Went Wrong"
		} else {
			utility.NewLogEntry(ctx).Info("201 - Created")
			response.Status = "Created"
			response.Code = http.StatusCreated
		}

	}
	// return response
	return ctx.JSON(response.Code, response)
}
