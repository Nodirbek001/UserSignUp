package service

import (
	"NewProUser/entity"
	"NewProUser/storage"
	"context"

	"github.com/docker/docker/daemon/logger"
)

type userServiceImpl struct{
	storage storage.Istorage
	logger logger.Logger
}

func NewUserService(log logger.Logger)UserService{
	return &userServiceImpl{
		storage: storage.NewStoragePg(),
		logger: log,
	}
}

type UserService interface{
	SignUpUser(ctx context.Context, req entity.SignUpModel)(string, error)
}

func(s *userServiceImpl) SignUpUser(cxt context.Context, model entity.SignUpModel)(string,error){
	return "", nil
}