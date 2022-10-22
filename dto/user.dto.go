package dto

type Register struct {
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique" validate:"required,email,unique"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}
