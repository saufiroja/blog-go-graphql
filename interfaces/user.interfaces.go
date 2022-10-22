package interfaces

import (
	"graphql/blog-go-graphql/dto"

	"github.com/graphql-go/graphql"
)

type UserRepository interface {
	Register(user *dto.Register) error
	Login(email string) (dto.Login, error)
}

type UserService interface {
	Register(user *dto.Register) error
	Login(email, password string) error
}

type UserResolvers interface {
	Register(params graphql.ResolveParams) (interface{}, error)
	Login(params graphql.ResolveParams) (interface{}, error)
}

type UserSchema interface {
	Mutation() *graphql.Object
	Query() *graphql.Object
	Root() *graphql.Schema
}
