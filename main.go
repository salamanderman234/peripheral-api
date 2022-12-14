package main

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/config"
	"github.com/salamanderman234/peripheral-api/middleware"
	"github.com/salamanderman234/peripheral-api/repository"
	utility "github.com/salamanderman234/peripheral-api/utility/log"
	"github.com/spf13/viper"
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
}

func main() {

	// make connection to database
	client, err := config.ConnectDB(config.GetDatabaseUri())
	if err != nil {
		utility.NewLogEntry(nil).Panic(err)
	}

	// dependency injection
	switchRepository := repository.NewSwitchRepository(client)

	// router configuration
	router := echo.New()
	router.HTTPErrorHandler = middleware.Error
	router.Use(middleware.Log)

	// router handler
}
