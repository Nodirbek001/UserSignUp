package postgres

import (
	"NewProUser/entity"
	"NewProUser/platforma/postgres"
	"NewProUser/storage/repo"
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"

)

type userRepo struct {
	db *gorm.DB
}

// func New(db *gorm.DB) *userRepo {
// 	return &userRepo{
// 		db: db,
// 	}
// }

func NewUserRepo() repo.UserStorageI {
	return &userRepo{
		db: postgres.DB(),
	}
}

func (r *userRepo) SignUpUser(ctx context.Context, model entity.SignUpModel) (string, error){
	err:=r.db.Table("users").WithContext(ctx).Create(model)
	if err!=nil {
		var pgErr *pgconn.PgError
		if errors.As(err.Error, &pgErr) && pgErr.Code=="23503" {
			return "",nil
		}
	}
	return model.ID, nil
}