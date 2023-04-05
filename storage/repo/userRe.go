package repo

import (
	"NewProUser/entity"
	"context"
)

type UserStorageI interface{
	SignUpUser(ctx context.Context, model entity.SignUpModel) (string, error)
}