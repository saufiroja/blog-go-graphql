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

	err := s.userRepository.Register(user)
	if err != nil {
		return errors.New("email already exists")
	}

	return nil
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.userRepository.Login(email)
	if err != nil {
		return "", errors.New("email is not registered")
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

func (s *UserService) FindUserById(id string) (entity.User, error) {
	return s.userRepository.FindUserById(id)
}

func (s *UserService) UpdateUser(id string, user *entity.User) error {
	return s.userRepository.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.userRepository.DeleteUser(id)
}
