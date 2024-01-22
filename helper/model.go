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
