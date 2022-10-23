package interfaces

import (
	"graphql/blog-go-graphql/dto"
	"graphql/blog-go-graphql/entity"

	"github.com/graphql-go/graphql"
)

type UserRepository interface {
	Register(user *dto.Register) error
	FindAllUsers() ([]entity.User, error)
	FindUserByID(id string) (entity.User, error)
}

type UserService interface {
	Register(user *dto.Register) error
	FindAllUsers() ([]entity.User, error)
	FindUserByID(id string) (entity.User, error)
}

type UserResolvers interface {
	Register(params graphql.ResolveParams) (interface{}, error)
	FindAllUsers(params graphql.ResolveParams) (interface{}, error)
	FindUserByID(params graphql.ResolveParams) (interface{}, error)
}

type UserSchema interface {
	Mutation() *graphql.Object
	Query() *graphql.Object
	Root() *graphql.Schema
}
