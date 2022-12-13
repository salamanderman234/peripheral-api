package main

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/middleware"
	"github.com/spf13/viper"
)

// membaca config json
func init() {
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func main() {
	router := echo.New()
	router.Use(middleware.Log)
}
