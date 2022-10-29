package service

import (
	"graphql/blog-go-graphql/database"
	"graphql/blog-go-graphql/dto"
	"graphql/blog-go-graphql/entity"
	"graphql/blog-go-graphql/interfaces"
)

type ArticleService struct {
	ArticleRepository interfaces.ArticleRepository
	Conf              database.Config
}

func NewArticleService(article interfaces.ArticleRepository, conf database.Config) interfaces.ArticleService {
	return &ArticleService{
		ArticleRepository: article,
		Conf:              conf,
	}
}

func (s *ArticleService) CreateArticle(article *dto.CreateArticle) error {
	err := s.ArticleRepository.CreateArticle(article)
	if err != nil {
		return err
	}

	return nil
}

func (s *ArticleService) FindAllArticles() ([]entity.Article, error) {
	return s.ArticleRepository.FindAllArticles()
}
