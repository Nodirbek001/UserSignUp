package entity

import (
	"NewProUser/utils"
	"fmt"
)

type SignUpModel struct {
	ID          string `json:"id" gorm:"id"`
	FullName    string `json:"fullname" gorm:"fullname"`
	PhoneNumber string `json:"phone_number" gorm:"phone_number"`
	Description string `json:"description" gorm:"description"`
	Photo       string `json:"photo" gorm:"photo"`
	Password    string `json:"password" gorm:"password"`
	Role        string `json:"role" gorm:"role"`
}

func Validate(phoneNumber, password string) error {
	if !utils.IsPhoenValid(phoneNumber) {
		return fmt.Errorf("invalid phone number")
	}
	if err := utils.ValidatePassword(password); err != nil {
		return err
	}
	return nil
}

type SignUpResModel struct{
	ID string `json:"id"`
	Access string `json:"access"`
	Refresh string `json:"refresh"`
}