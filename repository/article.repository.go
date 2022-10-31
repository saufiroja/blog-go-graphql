package repository

import (
	"graphql/blog-go-graphql/dto"
	"graphql/blog-go-graphql/entity"
	"graphql/blog-go-graphql/interfaces"

	"gorm.io/gorm"
)

type articleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(db *gorm.DB) interfaces.ArticleRepository {
	return &articleRepository{
		DB: db,
	}
}

func (r *articleRepository) CreateArticle(article *dto.CreateArticle) error {
	articleModel := entity.Article{
		PhotoURL: article.PhotoURL,
		Title:    article.Title,
		Body:     article.Body,
		Category: article.Category,
		UserID:   article.UserID,
	}

	err := r.DB.Model(&articleModel).Create(&articleModel).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *articleRepository) FindAllArticles() ([]entity.Article, error) {
	var articles []entity.Article

	err := r.DB.Model(&articles).Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *articleRepository) FindArticleByID(id string) (entity.Article, error) {
	var article entity.Article

	err := r.DB.Model(&article).Where("id = ?", id).Find(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}
