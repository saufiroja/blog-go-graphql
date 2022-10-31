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
	articleInput := &dto.CreateArticle{
		PhotoURL: params.Args["photoURL"].(string),
		Title:    params.Args["title"].(string),
		Body:     params.Args["body"].(string),
		Category: params.Args["category"].(string),
		UserID:   params.Args["userId"].(string),
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

func (r *articleResolver) FindArticleByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(string)

	article, err := r.articleService.FindArticleByID(id)
	if err != nil {
		return nil, err
	}

	return article, nil
}
