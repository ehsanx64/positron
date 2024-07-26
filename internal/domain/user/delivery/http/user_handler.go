package http

import (
	"net/http"

	"github.com/ehsanx64/positron/internal/domain/user/entity"
	"github.com/labstack/echo/v4"
)

type UserHTTPHandler struct {
}

func NewUserHTTPHandler(e *echo.Echo) *UserHTTPHandler {
	handler := &UserHTTPHandler{}

	e.GET("/users", handler.Fetch)
	return handler
}

func (uhh *UserHTTPHandler) Fetch(ctx echo.Context) error {
	var (
		users = []*entity.User{}
	)

	for i := 1; i < 5; i++ {
		users = append(users, entity.NewUser("User", i*10-5))
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Hello",
		"users":   users,
	})
}
