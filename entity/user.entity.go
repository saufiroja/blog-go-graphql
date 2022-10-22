package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique" validate:"required,email,unique"`
	Password  string    `json:"password" validate:"required,min=8,max=20"`
	Article   []Article `json:"article" gorm:"foreignKey:UserID"`
	CreatedAt int64     `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64     `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt int64     `json:"deleted_at" gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
