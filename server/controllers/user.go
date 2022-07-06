package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mococa/golang-api-boilerplate/server/dtos"
	"github.com/mococa/golang-api-boilerplate/server/services"
)

/* -------------- Types -------------- */
type UserController struct {
	services.UserService
}

/* -------------- Initialization -------------- */
func NewUserController(s *services.UserService) UserController {
	return UserController{*s}
}

/* -------------- Controllers -------------- */
func (controller *UserController) CreateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var user dtos.CreateUserDTO

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)

		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{"error": "Error unmarshalling"}`))
		return
	}

	created_user, err := controller.UserService.CreateUserService(&user)

	if err != nil {
		fmt.Println(err)

		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	response, err := json.Marshal(created_user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling"}`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(response))

}

func (controller *UserController) ListUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	users, err := controller.UserService.ListUserService()

	if err != nil {
		fmt.Println(err)

		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	response, err := json.Marshal(users)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling"}`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(response))

}
