package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	userHttpHandler "github.com/ehsanx64/positron/internal/domain/user/delivery/http"
)

func main() {
	e := echo.New()

	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("failed to load the config file")
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}

	// Config file found and successfully parsed

	userHttpHandler.NewUserHTTPHandler(e)
	e.Logger.Fatal(e.Start(viper.Get("app.port").(string)))
}
