package resolvers

import (
	"graphql/blog-go-graphql/dto"
	"graphql/blog-go-graphql/interfaces"
	"graphql/blog-go-graphql/utils"

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

	data := utils.ResponseSuccess{
		Code:    201,
		Message: "Register Success",
		Result:  user,
	}

	return data, nil
}

func (r *UserResolvers) Login(params graphql.ResolveParams) (interface{}, error) {
	email := params.Args["email"].(string)
	password := params.Args["password"].(string)
	token, err := r.UserService.Login(email, password)
	if err != nil {
		return nil, err
	}

	data := utils.ResponseSuccess{
		Code:    200,
		Message: "Login Success",
		Result: map[string]interface{}{
			"accessToken": token,
		},
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
