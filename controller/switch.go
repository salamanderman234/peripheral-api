package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/domain"
	"github.com/salamanderman234/peripheral-api/entity"
	"github.com/salamanderman234/peripheral-api/policy"
	utility "github.com/salamanderman234/peripheral-api/utility"
)

type switchController struct {
	service domain.SwitchService
}

func NewSwitchController(service domain.SwitchService) domain.SwitchController {
	return &switchController{
		service: service,
	}
}

func (s *switchController) GetOneSwitch(ctx echo.Context) error {
	var filter entity.Switch
	var foundSwitch []entity.Switch
	filter.SwitchID = ctx.Param("switch_id")

	// calling service
	foundSwitch, err := s.service.GetSwitch(ctx.Request().Context(), filter, "")
	if err != nil {
		go utility.NewLogEntry(ctx).Error("500 - Internal Server Error")
		return ctx.JSON(http.StatusInternalServerError, entity.BaseResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "Something went wrong",
		})
	}

	if len(foundSwitch) == 0 {
		go utility.NewLogEntry(ctx).Error("404 - Not Found")
		return ctx.JSON(http.StatusNotFound, entity.BaseResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Errors: "No matching data for the given query",
		})
	}
	// return response
	go utility.NewLogEntry(ctx).Info("200 - Ok")
	return ctx.JSON(http.StatusOK, entity.BaseResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   foundSwitch,
	})
}

func (s *switchController) GetAllSwitch(ctx echo.Context) error {
	// init
	var filter entity.Switch
	var foundSwitches []entity.Switch
	var sort string
	// get filter from query, body or path params
	ctx.Bind(&filter)
	sort = ctx.QueryParam("sort")
	// calling service
	foundSwitches, err := s.service.GetSwitch(ctx.Request().Context(), filter, sort)
	if err != nil {
		go utility.NewLogEntry(ctx).Error("500 - Internal Server Error")
		return ctx.JSON(http.StatusBadRequest, entity.BaseResponse{
			Status: "Internal Server Error",
			Code:   http.StatusInternalServerError,
			Errors: "Something Went Wrong",
		})
	}

	// check if result is empty
	if len(foundSwitches) == 0 {
		go utility.NewLogEntry(ctx).Error("404 - Not Found")
		return ctx.JSON(http.StatusBadRequest, entity.BaseResponse{
			Status: "Not Found",
			Code:   http.StatusNotFound,
			Errors: "No matching data for the given query",
		})
	}

	// sending response
	go utility.NewLogEntry(ctx).Info("200 - ok")
	return ctx.JSON(http.StatusOK, entity.BaseResponse{
		Status: "Ok",
		Code:   http.StatusOK,
		Data:   foundSwitches,
	})
}

func (s *switchController) CreateNewSwitch(ctx echo.Context) error {
	// init
	var body []entity.Switch
	var policyResults []*policy.SwitchPolicy

	// binding
	if err := ctx.Bind(&body); err != nil || len(body) == 0 {
		go utility.NewLogEntry(ctx).Error("400 - Bad Request")
		return ctx.JSON(http.StatusBadRequest, entity.BaseResponse{
			Status: "Bad Request",
			Code:   http.StatusBadRequest,
			Errors: "Data body does not match specifications",
		})
	}

	// checking input policy
	for _, element := range body {
		result := policy.DocumentSwitchPolicy(ctx.Request().Context(), element, s.service, "insert")
		if result != nil {
			policyResults = append(policyResults, result)
		}
	}
	if len(policyResults) != 0 {
		go utility.NewLogEntry(ctx).Error("400 - Bad Request")
		return ctx.JSON(http.StatusBadRequest, entity.BaseResponse{
			Status: "Bad Request",
			Code:   http.StatusBadRequest,
			Errors: policyResults,
		})
	}

	// calling service
	insertedIds, err := s.service.CreateSwitch(ctx.Request().Context(), body)
	// error while parsing body
	if err != nil {
		go utility.NewLogEntry(ctx).Error("500 - Internal Server Error")
		return ctx.JSON(http.StatusInternalServerError, entity.BaseResponse{
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
	var body entity.Switch
	var filter entity.Switch

	// get updatefield from body
	err := ctx.Bind(&body)
	if err != nil || ctx.Request().ContentLength == 0 {
		go utility.NewLogEntry(ctx).Error("400 - Bad Request")
		return ctx.JSON(http.StatusBadRequest, entity.BaseResponse{
			Status: "Bad Request",
			Code:   http.StatusBadRequest,
			Errors: "Data body does not match specifications",
		})
	}
	// checking policy
	policyCheckResult := policy.DocumentSwitchPolicy(ctx.Request().Context(), body, s.service, "update")
	if policyCheckResult != nil {
		go utility.NewLogEntry(ctx).Error("400 - Bad Request")
		return ctx.JSON(http.StatusBadRequest, entity.BaseResponse{
			Status: "Bad Request",
			Code:   http.StatusBadRequest,
			Errors: policyCheckResult,
		})
	}

	// calling service
	filter.SwitchID = ctx.Param("switch_id")
	modifiedDocument, err := s.service.UpdateSwitch(ctx.Request().Context(), body, filter)
	if err != nil {
		go utility.NewLogEntry(ctx).Error("500 - Internal Server Error")
		return ctx.JSON(http.StatusInternalServerError, entity.BaseResponse{
			Status: "Internal Server Error",
			Code:   http.StatusInternalServerError,
			Errors: "Something went wrong",
		})
	}
	// checking if any document are updated
	if modifiedDocument == 0 {
		go utility.NewLogEntry(ctx).Error("404 - Not Found")
		return ctx.JSON(http.StatusNotFound, entity.BaseResponse{
			Status: "Not Found",
			Code:   http.StatusNotFound,
			Errors: "No data found with that parameter",
		})
	}
	// sending response
	go utility.NewLogEntry(ctx).Info("200 - Ok")
	return ctx.JSON(http.StatusOK, entity.BaseResponse{
		Status: "Ok",
		Code:   http.StatusOK,
		Data:   fmt.Sprintf("%d Document Modified Successfully", modifiedDocument),
	})
}

func (s *switchController) DropSwitch(ctx echo.Context) error {
	filter := ctx.Param("switch_id")
	if filter == "" {
		go utility.NewLogEntry(ctx).Error("400 - Bad Request")
		return ctx.JSON(http.StatusBadRequest, entity.BaseResponse{
			Status: "Bad Request",
			Code:   http.StatusBadRequest,
			Errors: "Missing switch id parameter",
		})
	}
	deletedIDs, err := s.service.DeleteSwitch(ctx.Request().Context(), filter)
	if err != nil {
		go utility.NewLogEntry(ctx).Error("500 - Internal Server Error")
		return ctx.JSON(http.StatusInternalServerError, entity.BaseResponse{
			Status: "Internal Server Error",
			Code:   http.StatusInternalServerError,
			Errors: "Something went wrong",
		})
	}

	if deletedIDs == 0 {
		go utility.NewLogEntry(ctx).Error("404 - Not Found")
		return ctx.JSON(http.StatusNotFound, entity.BaseResponse{
			Status: "Not Found",
			Code:   http.StatusNotFound,
			Errors: "No data found with that parameter",
		})
	}
	// sending response
	go utility.NewLogEntry(ctx).Info("200 - Ok")
	return ctx.JSON(http.StatusOK, entity.BaseResponse{
		Status: "Ok",
		Code:   http.StatusOK,
		Data:   fmt.Sprintf("%d Document Deleted Successfully", deletedIDs),
	})
}
