package schema

import (
	"graphql/blog-go-graphql/interfaces"

	"github.com/graphql-go/graphql"
)

type Schema interface {
	Query() *graphql.Object
	Mutation() *graphql.Object
	Root() *graphql.Schema
}

type rootSchema struct {
	userSchema    interfaces.UserSchema
	articleSchema interfaces.ArticleSchema
}

func NewRootSchema(userSchema interfaces.UserSchema, articleSchema interfaces.ArticleSchema) Schema {
	return &rootSchema{
		userSchema:    userSchema,
		articleSchema: articleSchema,
	}
}

func (s *rootSchema) Query() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"FindAllUsers":    s.userSchema.FindAllUsers(),
			"FindUserById":    s.userSchema.FindUserById(),
			"FindAllArticles": s.articleSchema.FindAllArticles(),
			"FindArticleById": s.articleSchema.FindArticleByID(),
		},
	})
}

func (r *rootSchema) Mutation() *graphql.Object {
	object := graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"Register":      r.userSchema.Register(),
			"Login":         r.userSchema.Login(),
			"UpdateUser":    r.userSchema.UpdateUser(),
			"DeleteUser":    r.userSchema.DeleteUser(),
			"CreateArticle": r.articleSchema.CreateArticle(),
			"UpdateArticle": r.articleSchema.UpdateArticle(),
			"DeleteArticle": r.articleSchema.DeleteArticle(),
		},
	}

	return graphql.NewObject(object)
}

func (r *rootSchema) Root() *graphql.Schema {
	rootMutation := r.Mutation()
	rootQuery := r.Query()

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	return &schema
}
