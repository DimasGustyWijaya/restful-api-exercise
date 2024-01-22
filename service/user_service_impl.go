package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"restful-api/helper"
	"restful-api/model"
	"restful-api/model/web"
	"restful-api/repository"
)

type UserServiceImpl struct {
	DB       *sql.DB
	user     repository.UserRepository
	Validate *validator.Validate
}

func NewUserService(db *sql.DB, user repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		DB:       db,
		user:     user,
		Validate: validate,
	}
}

func (service UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := model.User{
		Name: request.Name,
	}

	user = service.user.Save(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := model.User{
		Name: request.Name,
	}

	user = service.user.Update(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service UserServiceImpl) Delete(ctx context.Context, userId int) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	errorDelete := service.user.Delete(ctx, tx, userId)
	helper.PanicIfError(errorDelete)

	return nil
}

func (service UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, errorId := service.user.FindById(ctx, tx, userId)
	helper.PanicIfError(errorId)

	return helper.ToUserResponse(result)
}
