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

type ProductServiceImpl struct {
	product  repository.ProductRepository
	validate *validator.Validate
	DB       *sql.DB
}

func NewProductService(product repository.ProductRepository, db *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		product:  product,
		DB:       db,
		validate: validate,
	}
}

func (p ProductServiceImpl) Save(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	err := p.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := p.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	prod := model.Product{
		Name: request.Name,
		Qty:  request.Qty,
	}

	prod = p.product.Save(ctx, tx, prod)

	return helper.ToProductResponse(prod)

}

func (p ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	err := p.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := p.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	prod := model.Product{
		Name: request.Name,
		Qty:  request.Qty,
	}

	prod = p.product.Update(ctx, tx, prod)

	return helper.ToProductResponse(prod)
}

func (p ProductServiceImpl) Delete(ctx context.Context, pruductId int) error {
	tx, err := p.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	errorDelete := p.product.Delete(ctx, tx, pruductId)
	helper.PanicIfError(errorDelete)

	return nil
}

func (p ProductServiceImpl) FindById(ctx context.Context, productId int) web.ProductResponse {

	tx, err := p.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := p.product.FindById(ctx, tx, productId)

	return helper.ToProductResponse(result)
}
