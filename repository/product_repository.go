package repository

import (
	"context"
	"database/sql"
	"restful-api/model"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product model.Product) model.Product
	Update(ctx context.Context, tx *sql.Tx, product model.Product) model.Product
	Delete(ctx context.Context, tx *sql.Tx, pruductId int) error
	FindById(ctx context.Context, tx *sql.Tx, productId int) (model.Product, error)
}
