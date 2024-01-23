package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"restful-api/app"
	"restful-api/controller"
	"restful-api/helper"
	"restful-api/middleware"
	"restful-api/repository"
	"restful-api/service"
)

func main() {
	db := app.GetConnection()
	validate := validator.New()

	prodRepository := repository.NewProductRepository()
	prodService := service.NewProductService(prodRepository, db, validate)
	prodControler := controller.NewProductController(prodService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository, validate)
	userControler := controller.NewUserController(userService)

	router := app.NewRouter(userControler, prodControler)

	server := http.Server{
		Handler: middleware.NewAuthMiddleware(router),
		Addr:    "localhost:4400",
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
