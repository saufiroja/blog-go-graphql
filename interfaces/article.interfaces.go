package interfaces

import (
	"graphql/blog-go-graphql/dto"
	"graphql/blog-go-graphql/entity"

	"github.com/graphql-go/graphql"
)

type ArticleRepository interface {
	CreateArticle(article *dto.CreateArticle) error
	FindAllArticles() ([]entity.Article, error)
}

type ArticleService interface {
	CreateArticle(article *dto.CreateArticle) error
	FindAllArticles() ([]entity.Article, error)
}

type ArticleResolvers interface {
	CreateArticle(params graphql.ResolveParams) (interface{}, error)
	FindAllArticles(params graphql.ResolveParams) (interface{}, error)
}

type ArticleSchema interface {
	Mutation() *graphql.Object
	Query() *graphql.Object
	Root() *graphql.Schema
}
