package interfaces

import (
	"graphql/blog-go-graphql/dto"
	"graphql/blog-go-graphql/entity"

	"github.com/graphql-go/graphql"
)

type UserRepository interface {
	Register(user *dto.Register) error
	Login(email string) (entity.User, error)
	FindAllUsers() ([]entity.User, error)
	FindUserById(id string) (entity.User, error)
	UpdateUser(id string, user *entity.User) error
	DeleteUser(id string) error
}

type UserService interface {
	Register(user *dto.Register) error
	Login(login *dto.Login) (dto.Token, error)
	FindAllUsers() ([]entity.User, error)
	FindUserById(id string) (entity.User, error)
	UpdateUser(id string, user *entity.User) error
	DeleteUser(id string) error
}

type UserResolvers interface {
	Register(params graphql.ResolveParams) (interface{}, error)
	Login(params graphql.ResolveParams) (interface{}, error)
	FindAllUsers(params graphql.ResolveParams) (interface{}, error)
	FindUserById(params graphql.ResolveParams) (interface{}, error)
	UpdateUser(params graphql.ResolveParams) (interface{}, error)
	DeleteUser(params graphql.ResolveParams) (interface{}, error)
}

type UserSchema interface {
	Register() *graphql.Field
	Login() *graphql.Field
	FindAllUsers() *graphql.Field
	FindUserById() *graphql.Field
	UpdateUser() *graphql.Field
	DeleteUser() *graphql.Field
}
