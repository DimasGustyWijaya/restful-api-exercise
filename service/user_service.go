package service

import (
	"context"
	"restful-api/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, userId int) error
	FindById(ctx context.Context, userId int) web.UserResponse
}
