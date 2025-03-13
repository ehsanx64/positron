package http

import (
	"net/http"
)

type UserHTTPHandler struct{}

type UserResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewUserHTTPHandler(mux *http.ServeMux) {
	handler := &UserHTTPHandler{}
	mux.HandleFunc("/users", handler.Fetch)
}

func (uhh *UserHTTPHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	var (
		users = []UserResponse{}
	)

	for i := 1; i < 5; i++ {
		users = append(users, UserResponse{
			Name: "User",
			Age:  i*10 - 5,
		})
	}

	WriteJson(map[string]any{
		"message": "Current Users List",
		"users":   users,
	}, w)
}
