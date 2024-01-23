package repository

import (
	"context"
	"database/sql"
	"errors"
	"restful-api/helper"
	"restful-api/model"
)

type ProuctRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProuctRepositoryImpl{}
}

func (p ProuctRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product model.Product) model.Product {
	SQL := "INSERT INTO product(name,qty) VALUES (?,?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Qty)
	helper.PanicIfError(err)

	id, erre := result.LastInsertId()
	helper.PanicIfError(erre)

	product.IdProduct = int(id)

	return product
}

func (p ProuctRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product model.Product) model.Product {
	SQL := "UPDATE product SET name = (?) WHERE id = (?)"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.IdProduct)
	helper.PanicIfError(err)

	return product
}

func (p ProuctRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, pruductId int) error {
	SQL := "DELETE FROM product WHERE id = (?)"
	_, err := tx.ExecContext(ctx, SQL, pruductId)
	helper.PanicIfError(err)

	return nil
}

func (p ProuctRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (model.Product, error) {
	SQL := "SELECT id,name FROM user WHERE id = (?)"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	product := model.Product{}
	if rows.Next() {
		err := rows.Scan(&product.IdProduct, &product.Name)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("user not found")
	}

}
