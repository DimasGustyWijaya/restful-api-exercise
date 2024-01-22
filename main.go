package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"restful-api/app"
	"restful-api/controller"
	"restful-api/helper"
	"restful-api/repository"
	"restful-api/service"
)

func main() {
	db := app.GetConnection()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository, validate)
	userControler := controller.NewUserController(userService)

	router := app.NewRouter(userControler)

	server := http.Server{
		Handler: router,
		Addr:    "localhost:4400",
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
