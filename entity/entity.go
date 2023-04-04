package entity

type SignUpModel struct {
	ID             string `json:"id" gorm:"id"`
	FullName       string `json:"fullname" gorm:"fullname"`
	PhoneNumber    string `json:"phone_number" gorm:"phone_number"`
	Description    string `json:"description" gorm:"description"`
	Photo          string `json:"photo" gorm:"photo"`
}

