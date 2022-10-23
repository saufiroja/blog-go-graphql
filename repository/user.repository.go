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

func (r *userRepository) Login(email string) (entity.User, error) {
	user := entity.User{}

	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (r *userRepository) FindAllUsers() ([]entity.User, error) {
	var users []entity.User

	_ = r.DB.Find(&users).Error

	return users, nil
}

func (r *userRepository) FindUserByID(id string) (entity.User, error) {
	user := entity.User{}

	result := r.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
