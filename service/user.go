package service

import (
	"NewProUser/entity"
	"NewProUser/storage"
	"context"

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
	id, err:=s.storage.UserF().SignUpUser(cxt, model)
	if err!=nil {
	}
	return "", nil
}