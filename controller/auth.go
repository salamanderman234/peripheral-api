package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/domain"
	"github.com/salamanderman234/peripheral-api/entity"
	utility "github.com/salamanderman234/peripheral-api/utility"
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
	var creds entity.Credentials
	err := ctx.Bind(&creds)
	if err != nil || ctx.Request().ContentLength == 0 || (creds.Username == "" || creds.Password == "") {
		go utility.NewLogEntry(ctx).Error("400 - Bad Request")
		return ctx.JSON(http.StatusBadRequest, entity.BaseResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: "Data body does not match specifications",
		})
	}
	token, err := a.service.Authenticate(ctx.Request().Context(), creds)
	if err != nil {
		go utility.NewLogEntry(ctx).Error("500 - Internal Server Error")
		return ctx.JSON(http.StatusBadRequest, entity.BaseResponse{
			Status: "Internal Server Error",
			Code:   http.StatusInternalServerError,
			Errors: "Something Went Wrong",
		})
	}
	if token == "" {
		go utility.NewLogEntry(ctx).Error("401 - Unauthorized")
		return ctx.JSON(http.StatusUnauthorized, entity.BaseResponse{
			Status: "Unauthorized",
			Code:   http.StatusUnauthorized,
			Errors: "Wrong credentials",
		})
	}
	// creating cookie
	cookie := http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(5 * time.Hour),
		Path:    "/",
	}
	ctx.SetCookie(&cookie)
	go utility.NewLogEntry(ctx).Info("200 - Ok")
	return ctx.JSON(http.StatusOK, entity.BaseResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   "Autenticate successfully",
	})
}
func (a *authController) Register(ctx echo.Context) error {
	return nil
}
