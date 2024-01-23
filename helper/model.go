package helper

import (
	"restful-api/model"
	"restful-api/model/web"
)

func ToUserResponse(user model.User) web.UserResponse {
	return web.UserResponse{
		Id:   user.Id,
		Name: user.Name,
	}
}

func ToProductResponse(prod model.Product) web.ProductResponse {
	return web.ProductResponse{
		IdProduct: prod.IdProduct,
		Name:      prod.Name,
		Qty:       prod.Qty,
	}
}
