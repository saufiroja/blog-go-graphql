package repository

import (
	"graphql/blog-go-graphql/dto"
	"graphql/blog-go-graphql/entity"
	"graphql/blog-go-graphql/interfaces"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) Register(user *dto.Register) error {
	users := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	result := r.DB.Model(&users).Create(&users)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
