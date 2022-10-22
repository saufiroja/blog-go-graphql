package service

import (
	"graphql/blog-go-graphql/database"
	"graphql/blog-go-graphql/dto"
	"graphql/blog-go-graphql/interfaces"
	"graphql/blog-go-graphql/utils"
)

type UserService struct {
	userRepository interfaces.UserRepository
	conf           database.Config
}

func NewUserService(userRepository interfaces.UserRepository, conf database.Config) interfaces.UserService {
	return &UserService{
		userRepository: userRepository,
		conf:           conf,
	}
}

func (s *UserService) Register(user *dto.Register) error {
	hash, _ := utils.HashPassword(user.Password)
	user.Password = hash

	return s.userRepository.Register(user)
}

func (s *UserService) Login(email, password string) error {
	user, err := s.userRepository.Login(email)
	if err != nil {
		return err
	}

	err = utils.ComparePassword(user.Password, password)
	if err != nil {
		return err
	}

	return nil
}
