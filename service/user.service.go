package service

import (
	"errors"
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

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.userRepository.Login(email)
	if err != nil {
		return "", err
	}

	err = utils.ComparePassword(user.Password, password)
	if err != nil {
		return "", errors.New("password is incorrect")
	}

	token, _ := utils.GenerateAccessToken(user.ID, user.Email)

	return token, nil
}

func (s *UserService) FindAllUsers() ([]entity.User, error) {
	return s.userRepository.FindAllUsers()
}
