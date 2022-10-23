package schema

import (
	"graphql/blog-go-graphql/interfaces"

	"github.com/graphql-go/graphql"
)

type UserSchema struct {
	UserResolvers interfaces.UserResolvers
}

func NewUserSchema(userResolvers interfaces.UserResolvers) interfaces.UserSchema {
	return &UserSchema{
		UserResolvers: userResolvers,
	}
}

// variable for user schema
var (
	User = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "User Schema",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.Int,
			},
			"updated_at": &graphql.Field{
				Type: graphql.Int,
			},
			"deleted_at": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})
)

// Query for user schema
func (s *UserSchema) Query() *graphql.Object {
	object := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"FindAllUsers": &graphql.Field{
				Type:        graphql.NewList(User),
				Description: "Get All Users",
				Resolve:     s.UserResolvers.FindAllUsers,
			},
		},
	}

	return graphql.NewObject(object)
}

// Mutation for user schema
func (s *UserSchema) Mutation() *graphql.Object {
	object := graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"Register": &graphql.Field{
				Type:        User,
				Description: "Register User",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: s.UserResolvers.Register,
			},
			"Login": &graphql.Field{
				Type:        User,
				Description: "Login User",
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
			},
		},
	}

	return graphql.NewObject(object)
}

// root mutation and query
func (s *UserSchema) Root() *graphql.Schema {
	root := graphql.SchemaConfig{
		Query:    s.Query(),
		Mutation: s.Mutation(),
	}

	schema, err := graphql.NewSchema(root)
	if err != nil {
		panic(err)
	}

	return &schema
}
