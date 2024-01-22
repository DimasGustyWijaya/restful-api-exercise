package app

import (
	"github.com/julienschmidt/httprouter"
	"restful-api/controller"
)

func NewRouter(controller controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/user", controller.Create)
	router.PUT("/api/user/:userId", controller.Update)
	router.DELETE("/api/user/:userId", controller.Delete)
	router.GET("/api/user/:userId", controller.FindById)

	return router

}
