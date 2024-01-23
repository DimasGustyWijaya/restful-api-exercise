package repository

import (
	"context"
	"database/sql"
	"errors"
	"restful-api/helper"
	"restful-api/model"
	"sync"
)

type ProuctRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProuctRepositoryImpl{}
}

func (p ProuctRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product model.Product) model.Product {
	SQL := "INSERT INTO product(name,qty) VALUES (?,?)"

	locker := sync.Mutex{}

	locker.Lock()
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Qty)
	locker.Unlock()

	helper.PanicIfError(err)

	id, erre := result.LastInsertId()
	helper.PanicIfError(erre)

	product.IdProduct = int(id)

	return product
}

func (p ProuctRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product model.Product) model.Product {
	SQL := "UPDATE product SET name = (?) WHERE id = (?)"

	locker := sync.Mutex{}

	locker.Lock()
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.IdProduct)
	locker.Unlock()

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
	SQL := "SELECT idproduct,name,qty FROM  WHERE id = (?)"

	lockRead := sync.RWMutex{}

	lockRead.RLock()
	rows, err := tx.QueryContext(ctx, SQL, productId)
	lockRead.RUnlock()

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
