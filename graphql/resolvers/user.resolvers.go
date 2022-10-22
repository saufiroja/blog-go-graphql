package resolvers

import (
	"graphql/blog-go-graphql/dto"
	"graphql/blog-go-graphql/entity"
	"graphql/blog-go-graphql/interfaces"

	"github.com/graphql-go/graphql"
)

type UserResolvers struct {
	UserService interfaces.UserService
}

func NewUserResolvers(userService interfaces.UserService) interfaces.UserResolvers {
	return &UserResolvers{
		UserService: userService,
	}
}

func (r *UserResolvers) Register(params graphql.ResolveParams) (interface{}, error) {
	user := &dto.Register{
		Name:     params.Args["name"].(string),
		Email:    params.Args["email"].(string),
		Password: params.Args["password"].(string),
	}

	err := r.UserService.Register(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserResolvers) Login(params graphql.ResolveParams) (interface{}, error) {
	user := &entity.User{
		Email:    params.Args["email"].(string),
		Password: params.Args["password"].(string),
	}

	err := r.UserService.Login(user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserResolvers) FindAllUsers(params graphql.ResolveParams) (interface{}, error) {
	users, err := r.UserService.FindAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
