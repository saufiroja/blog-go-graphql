package schema

import (
	"graphql/blog-go-graphql/interfaces"

	"github.com/graphql-go/graphql"
)

type userSchema struct {
	userResolvers interfaces.UserResolvers
}

func NewUserSchema(userResolvers interfaces.UserResolvers) interfaces.UserSchema {
	return &userSchema{
		userResolvers: userResolvers,
	}
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"role": &graphql.Field{
			Type: graphql.String,
		},
		"craetedAt": &graphql.Field{
			Type: graphql.Int,
		},
		"updatedAt": &graphql.Field{
			Type: graphql.Int,
		},
		"deletedAt": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var loginType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Login",
	Fields: graphql.Fields{
		"accessToken": &graphql.Field{
			Type: graphql.String,
		},
		"refreshToken": &graphql.Field{
			Type: graphql.String,
		},
		"expiresIn": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
})

func (u *userSchema) Register() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"role": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	field := graphql.Field{
		Type:        userType,
		Description: "Register a new user",
		Args:        args,
		Resolve:     u.userResolvers.Register,
	}

	return &field
}

func (u *userSchema) Login() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	field := graphql.Field{
		Type:        loginType,
		Description: "Login a user",
		Args:        args,
		Resolve:     u.userResolvers.Login,
	}

	return &field
}

func (u *userSchema) FindAllUsers() *graphql.Field {
	field := graphql.Field{
		Type:        graphql.NewList(userType),
		Description: "Find all users",
		Resolve:     u.userResolvers.FindAllUsers,
	}

	return &field
}

func (u *userSchema) FindUserById() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	field := graphql.Field{
		Type:        userType,
		Description: "Find a user by id",
		Args:        args,
		Resolve:     u.userResolvers.FindUserById,
	}

	return &field
}

func (u *userSchema) UpdateUser() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"input": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.NewInputObject(graphql.InputObjectConfig{
				Name: "UpdateUserInput",
				Fields: graphql.InputObjectConfigFieldMap{
					"name": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
				},
			})),
		},
	}

	field := graphql.Field{
		Type:        userType,
		Description: "Update a user",
		Args:        args,
		Resolve:     u.userResolvers.UpdateUser,
	}

	return &field
}

func (u *userSchema) DeleteUser() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	field := graphql.Field{
		Type:        userType,
		Description: "Delete a user",
		Args:        args,
		Resolve:     u.userResolvers.DeleteUser,
	}

	return &field
}
