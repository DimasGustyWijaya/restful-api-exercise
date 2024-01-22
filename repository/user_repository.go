package repository

import (
	"context"
	"database/sql"
	"restful-api/model"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user model.User) model.User
	Update(ctx context.Context, tx *sql.Tx, user model.User) model.User
	Delete(ctx context.Context, tx *sql.Tx, userId int) error
	FindById(ctx context.Context, tx *sql.Tx, userId int) (model.User, error)
}
