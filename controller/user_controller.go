package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bagasjs/go-blog/model"
	"github.com/bagasjs/go-blog/service"
	"github.com/gorilla/mux"
)

type UserController struct {
    userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
    return &UserController{
        userService: service,
    }
}

func (controller *UserController) Route(r *mux.Router) {
    r.HandleFunc("/api/users", controller.allUsers).Methods("GET")
    r.HandleFunc("/api/users", controller.createUser).Methods("POST")
}

func (controller *UserController) allUsers(w http.ResponseWriter, r *http.Request) {
    data, err := controller.userService.List()
    response := model.WebResponse{ Status: "All Users Fetched", Code: 200, Data: data }
    if err != nil {
        response.Code = 400
        response.Status = fmt.Sprint("Failed to fetch all users: ", err)
    }

    responseBytes, err := json.Marshal(response)
    if err != nil {
        r.Response.StatusCode = 500
        fmt.Fprint(w, "Failed to encode data in JSON format")
    }
    w.Header().Add("Content-Type", "application/json")
    fmt.Fprint(w, string(responseBytes))
}

func (controller *UserController) createUser(w http.ResponseWriter, r *http.Request) {
    createUserRequest := model.CreateUserRequest {
        Name: r.FormValue("name"),
        Email: r.FormValue("email"),
        Password: r.FormValue("password"),
        PasswordConfirmation: r.FormValue("password_confirmation"),
    }
    createUserResponse, err := controller.userService.Create(createUserRequest)
    response := model.WebResponse{ Status: "User generated", Code: 200, Data: createUserResponse }
    if err != nil {
        response.Code = 400
        response.Status = fmt.Sprint("Failed to generate user: ", err)
    }

    responseBytes, err := json.Marshal(response)
    if err != nil {
        r.Response.StatusCode = 500
        fmt.Fprint(w, "Failed to encode data in JSON format")
    }

    w.Header().Add("Content-Type", "application/json")
    fmt.Fprint(w, string(responseBytes))
}
