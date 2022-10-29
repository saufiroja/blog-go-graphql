package schema

import (
	"graphql/blog-go-graphql/interfaces"

	"github.com/graphql-go/graphql"
)

type ArticleSchema struct {
	articleResolvers interfaces.ArticleResolvers
}

func NewArticleSchema(articleResolvers interfaces.ArticleResolvers) interfaces.ArticleSchema {
	return &ArticleSchema{
		articleResolvers: articleResolvers,
	}
}

// variable for article schema
var (
	Article = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Article",
		Description: "Article",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"photoURL": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"body": &graphql.Field{
				Type: graphql.String,
			},
			"category": &graphql.Field{
				Type: graphql.String,
			},
			"userID": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.Int,
			},
			"updated_at": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})
)

func (s *ArticleSchema) Query() *graphql.Object {
	object := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"articles": &graphql.Field{
				Type:        graphql.NewList(Article),
				Description: "Get all articles",
				Resolve:     s.articleResolvers.FindAllArticles,
			},
		},
	}

	return graphql.NewObject(object)
}

func (s *ArticleSchema) Mutation() *graphql.Object {
	object := graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createArticle": &graphql.Field{
				Type:        Article,
				Description: "Create new article",
				Args: graphql.FieldConfigArgument{
					"photoURL": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"body": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"userID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: s.articleResolvers.CreateArticle,
			},
		},
	}

	return graphql.NewObject(object)
}

func (s *ArticleSchema) Root() *graphql.Schema {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    s.Query(),
		Mutation: s.Mutation(),
	})

	return &schema
}
