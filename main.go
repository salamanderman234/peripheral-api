package main

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/config"
	"github.com/salamanderman234/peripheral-api/controller"
	"github.com/salamanderman234/peripheral-api/domain"
	"github.com/salamanderman234/peripheral-api/middleware"
	"github.com/salamanderman234/peripheral-api/repository"
	route "github.com/salamanderman234/peripheral-api/routes"
	"github.com/salamanderman234/peripheral-api/service"
	utility "github.com/salamanderman234/peripheral-api/utility"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// membaca config json
func init() {
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		utility.NewLogEntry(nil).Panic(err)
	}

	logrus.SetFormatter(&prefixed.TextFormatter{})
}

func main() {

	// make connection to database
	client, err := config.ConnectDB(config.GetDatabaseUri())
	if err != nil {
		utility.NewLogEntry(nil).Panic(err)
	}
	// dependency injection
	switchRepository := repository.NewSwitchRepository(client)
	authRepository := repository.NewAuthRepository(client)

	switchService := service.NewSwitchService(switchRepository)
	authService := service.NewAuthService(authRepository)

	switchController := controller.NewSwitchController(switchService)
	authController := controller.NewAuthController(authService)

	// router configuration
	router := echo.New()
	router.HTTPErrorHandler = middleware.Error
	router.Use(middleware.Log)

	// router handler
	group := router.Group(config.CreateAppPath(""))
	routersList := []domain.Router{
		route.NewSwitchRoute(group, switchController),
		route.NewAuthRoute(group, authController),
	}
	// /api/v1/
	for _, router := range routersList {
		router.RegisterRoutes()
	}

	router.Logger.Fatal(router.Start(viper.GetString("app.port")))
}
