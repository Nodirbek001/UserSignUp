package storage

import (
	"NewProUser/storage/postgres"
	"NewProUser/storage/repo"
)

type storagePg struct {
	userRepo repo.UserStorageI
}

type Istorage interface {
	UserF() repo.UserStorageI
}

func (s storagePg) UserF() repo.UserStorageI {
	return s.userRepo
}

func NewStoragePg() Istorage {
	return &storagePg{userRepo: postgres.NewUserRepo()}
}
