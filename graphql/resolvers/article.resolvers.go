package resolvers

import (
	"graphql/blog-go-graphql/dto"
	"graphql/blog-go-graphql/interfaces"

	"github.com/graphql-go/graphql"
)

type articleResolver struct {
	articleService interfaces.ArticleService
}

func NewArticleResolvers(articleService interfaces.ArticleService) interfaces.ArticleResolvers {
	return &articleResolver{
		articleService: articleService,
	}
}

func (r *articleResolver) CreateArticle(params graphql.ResolveParams) (interface{}, error) {
	article := params.Args["article"].(map[string]interface{})

	articleInput := &dto.CreateArticle{
		PhotoURL: article["photoURL"].(string),
		Title:    article["title"].(string),
		Body:     article["body"].(string),
		Category: article["category"].(string),
		UserID:   article["userID"].(string),
	}

	err := r.articleService.CreateArticle(articleInput)
	if err != nil {
		return nil, err
	}

	return articleInput, nil
}

func (r *articleResolver) FindAllArticles(params graphql.ResolveParams) (interface{}, error) {
	articles, err := r.articleService.FindAllArticles()
	if err != nil {
		return nil, err
	}

	return articles, nil
}
