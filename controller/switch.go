package controller

import (
	"encoding/json"
	"fmt"
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

func (s *switchController) GetOneSwitch(ctx echo.Context) error {
	var filter entity.Switch
	var foundSwitch []entity.Switch
	filter.Slug = ctx.Param("slug")

	// calling service
	result, err := s.service.GetSwitch(ctx.Request().Context(), filter)
	if err != nil {
		go utility.NewLogEntry(ctx).Error("500 - Internal Server Error")
		return ctx.JSON(http.StatusInternalServerError, entity.BaseResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "Something went wrong",
		})
	}

	// decode
	json.Unmarshal(result, &foundSwitch)

	// return response
	go utility.NewLogEntry(ctx).Error("200 - Ok")
	return ctx.JSON(http.StatusOK, entity.BaseResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   foundSwitch,
	})
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
		go utility.NewLogEntry(ctx).Error("500 - Internal Server Error")
		return ctx.JSON(http.StatusBadRequest, entity.BaseResponse{
			Status: "Internal Server Error",
			Code:   http.StatusInternalServerError,
			Errors: "Something Went Wrong",
		})
	}
	// convert result to array of entitiy.switch
	json.Unmarshal(result, &switchs)

	if len(switchs) == 0 {
		go utility.NewLogEntry(ctx).Error("404 - Not Found")
		return ctx.JSON(http.StatusBadRequest, entity.BaseResponse{
			Status: "Not Found",
			Code:   http.StatusNotFound,
			Errors: "No matching data for the given query",
		})
	}

	// sending response
	go utility.NewLogEntry(ctx).Info("200 - ok")
	return ctx.JSON(response.Code, entity.BaseResponse{
		Status: "Ok",
		Code:   http.StatusOK,
		Data:   switchs,
	})
}

func (s *switchController) CreateNewSwitch(ctx echo.Context) error {
	// init
	var switchsBody []entity.Switch

	// binding
	if err := ctx.Bind(&switchsBody); err != nil {
		go utility.NewLogEntry(ctx).Error("400 - Bad Request")
		return ctx.JSON(http.StatusBadRequest, entity.BaseResponse{
			Status: "Bad Request",
			Code:   http.StatusBadRequest,
			Errors: "Data body does not match specifications",
		})
	}

	// calling service
	insertedIds, policy, err := s.service.CreateSwitch(ctx.Request().Context(), switchsBody)
	// checking policy
	if policy != nil {
		go utility.NewLogEntry(ctx).Error("400 - Bad Request")
		return ctx.JSON(http.StatusBadRequest, entity.BaseResponse{
			Status: "Bad Request",
			Code:   http.StatusBadRequest,
			Errors: policy,
		})
	}
	// error while parsing body
	if err != nil {
		go utility.NewLogEntry(ctx).Error("500 - Internal Server Error")
		return ctx.JSON(http.StatusBadRequest, entity.BaseResponse{
			Status: "Internal Server Error",
			Code:   http.StatusInternalServerError,
			Errors: "Something Went Wrong",
		})
	}

	// return response
	go utility.NewLogEntry(ctx).Info("201 - Created")
	return ctx.JSON(http.StatusCreated, entity.BaseResponse{
		Status: "Created",
		Code:   http.StatusCreated,
		Data:   fmt.Sprintf("%d Document Inserted Successfully", len(insertedIds)),
	})
}

func (s *switchController) UpdateOneSwitch(ctx echo.Context) error {
	// init
	var updateField entity.Switch
	var filter entity.Switch

	// get updatefield from body
	err := ctx.Bind(&updateField)
	if err != nil {
		go utility.NewLogEntry(ctx).Error("400 - Bad Request")
		return ctx.JSON(http.StatusBadRequest, entity.BaseResponse{
			Status: "Bad Request",
			Code:   http.StatusBadRequest,
			Errors: "Data body does not match specifications",
		})
	}
	// var foundSwitch []entity.Switch
	filter.Slug = ctx.Param("slug")
	// calling service
	modifiedDocument, err := s.service.UpdateSwitch(ctx.Request().Context(), updateField, filter)
	if err != nil {
		go utility.NewLogEntry(ctx).Error("500 - Internal Server Error")
		return ctx.JSON(http.StatusInternalServerError, entity.BaseResponse{
			Status: "Internal Server Error",
			Code:   http.StatusInternalServerError,
			Errors: "Something went wrong",
		})
	}

	if modifiedDocument == 0 {
		go utility.NewLogEntry(ctx).Error("404 - Not Found")
		return ctx.JSON(http.StatusNotFound, entity.BaseResponse{
			Status: "Not Found",
			Code:   http.StatusNotFound,
			Errors: "No data found with that parameter",
		})
	}

	go utility.NewLogEntry(ctx).Info("200 - Ok")
	return ctx.JSON(http.StatusOK, entity.BaseResponse{
		Status: "Ok",
		Code:   http.StatusOK,
		Data:   fmt.Sprintf("%d Document Modified Successfully", modifiedDocument),
	})
}
