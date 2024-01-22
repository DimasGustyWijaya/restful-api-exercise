package repository

import (
	"context"
	"database/sql"
	"errors"
	"restful-api/helper"
	"restful-api/model"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	SQL := "INSERT INTO user(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, user.Name)
	helper.PanicIfError(err)

	id, erre := result.LastInsertId()
	helper.PanicIfError(erre)

	user.Id = int(id)

	return user
}

func (u *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	SQL := "UPDATE user SET name = (?) WHERE id = (?)"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Id)
	helper.PanicIfError(err)

	return user
}

func (u *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userId int) error {
	SQL := "DELETE FROM user WHERE id = (?)"
	_, err := tx.ExecContext(ctx, SQL, userId)
	helper.PanicIfError(err)

	return nil
}

func (u *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (model.User, error) {
	SQL := "SELECT id,name FROM user WHERE id = (?)"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := model.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}

}
