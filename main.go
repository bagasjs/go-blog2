package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/bagasjs/go-blog/controller"
	"github.com/bagasjs/go-blog/repository"
	"github.com/bagasjs/go-blog/service"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()

    userRepository := repository.NewUserSQLite3Repository(db)
    userService := service.NewUserService(userRepository)
    userController := controller.NewUserController(userService)
    userController.Route(r)

    log.Println("Server is running at port 8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
