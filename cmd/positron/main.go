package main

import (
	"fmt"

	userHttpHandler "github.com/ehsanx64/positron/internal/domain/user/delivery/http"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello world!!!")
	e := echo.New()

	userHttpHandler.NewUserHTTPHandler(e)
	e.Logger.Fatal(e.Start(":1323"))
}
