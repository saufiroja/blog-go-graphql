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

func (r *userRepository) Login(email string) (dto.Login, error) {
	login := dto.Login{
		Email: email,
	}
	user := entity.User{
		Email: login.Email,
	}

	result := r.DB.Model(&user).Where("email = ?", login.Email).First(&user)
	if result.Error != nil {
		return login, result.Error
	}

	return login, nil
}

func (r *userRepository) FindAllUsers() ([]entity.User, error) {
	var users []entity.User

	_ = r.DB.Find(&users).Error

	return users, nil
}
