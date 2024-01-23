package app

import (
	"github.com/julienschmidt/httprouter"
	"restful-api/controller"
)

func NewRouter(user controller.UserController, product controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	// User Handler
	router.POST("/api/user", user.Create)
	router.PUT("/api/user/:userId", user.Update)
	router.DELETE("/api/user/:userId", user.Delete)
	router.GET("/api/user/:userId", user.FindById)

	// Product Handler
	router.POST("/api/product", product.Create)
	router.PUT("/api/product/:productId", product.Update)

	return router

}
