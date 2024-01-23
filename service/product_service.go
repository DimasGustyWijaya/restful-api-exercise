package service

import (
	"context"
	"restful-api/model/web"
)

type ProductService interface {
	Save(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	Delete(ctx context.Context, pruductId int) error
	FindById(ctx context.Context, productId int) web.ProductResponse
}
