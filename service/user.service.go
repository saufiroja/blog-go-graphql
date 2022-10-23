package service

import (
	"graphql/blog-go-graphql/dto"
	"graphql/blog-go-graphql/entity"
	"graphql/blog-go-graphql/interfaces"
	"graphql/blog-go-graphql/utils"
)

type UserService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) interfaces.UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Register(user *dto.Register) error {
	hash, _ := utils.HashPassword(user.Password)
	user.Password = hash

	return s.userRepository.Register(user)
}

func (s *UserService) FindAllUsers() ([]entity.User, error) {
	return s.userRepository.FindAllUsers()
}

func (s *UserService) FindUserByID(id string) (entity.User, error) {
	return s.userRepository.FindUserByID(id)
}
