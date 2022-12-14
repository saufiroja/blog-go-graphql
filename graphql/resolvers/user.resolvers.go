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
		Role:     params.Args["role"].(string),
	}

	err := r.UserService.Register(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserResolvers) Login(params graphql.ResolveParams) (interface{}, error) {
	user := &dto.Login{
		Email:    params.Args["email"].(string),
		Password: params.Args["password"].(string),
	}
	token, err := r.UserService.Login(user)
	if err != nil {
		return nil, err
	}

	data := dto.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresIn,
	}

	return data, nil
}

func (r *UserResolvers) FindAllUsers(params graphql.ResolveParams) (interface{}, error) {
	users, err := r.UserService.FindAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserResolvers) FindUserById(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(string)
	user, err := r.UserService.FindUserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserResolvers) UpdateUser(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(string)
	user := &entity.User{
		Name: params.Args["name"].(string),
	}

	err := r.UserService.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserResolvers) DeleteUser(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(string)
	err := r.UserService.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
