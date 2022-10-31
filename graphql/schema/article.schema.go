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

var articleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Article",
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
		"userId": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
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

func (s *ArticleSchema) CreateArticle() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"photoURL": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"body": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"category": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"userId": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	field := graphql.Field{
		Type:        articleType,
		Description: "Create a new article",
		Args:        args,
		Resolve:     s.articleResolvers.CreateArticle,
	}

	return &field
}

func (s *ArticleSchema) FindAllArticles() *graphql.Field {
	field := graphql.Field{
		Type:        graphql.NewList(articleType),
		Description: "Find all articles",
		Resolve:     s.articleResolvers.FindAllArticles,
	}

	return &field
}

func (s *ArticleSchema) FindArticleByID() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	field := graphql.Field{
		Type:        articleType,
		Description: "Find article by ID",
		Args:        args,
		Resolve:     s.articleResolvers.FindArticleByID,
	}

	return &field
}

func (s *ArticleSchema) UpdateArticle() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"photoURL": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"body": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"category": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"userId": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	field := graphql.Field{
		Type:        articleType,
		Description: "Update an article",
		Args:        args,
		Resolve:     s.articleResolvers.UpdateArticle,
	}

	return &field
}

func (s *ArticleSchema) DeleteArticle() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	field := graphql.Field{
		Type:        articleType,
		Description: "Delete an article",
		Args:        args,
		Resolve:     s.articleResolvers.DeleteArticle,
	}

	return &field
}
