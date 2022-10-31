package service

import (
	"errors"
	"graphql/blog-go-graphql/database"
	"graphql/blog-go-graphql/dto"
	"graphql/blog-go-graphql/entity"
	"graphql/blog-go-graphql/interfaces"
	"sync"
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

func (s *ArticleService) FindArticleByID(id string) (entity.Article, error) {
	return s.ArticleRepository.FindArticleByID(id)
}

func (s *ArticleService) UpdateArticle(id string, article *dto.UpdateArticle) error {
	wg := sync.WaitGroup{}
	var err error
	wg.Add(2)

	go func() {
		defer wg.Done()
		_, err = s.ArticleRepository.FindArticleByID(id)
		if err != nil {
			err = errors.New("article not found")
		}
	}()

	go func() {
		defer wg.Done()
		err = s.ArticleRepository.UpdateArticle(id, article)
		if err != nil {
			err = errors.New("failed to update article")
		}
	}()

	wg.Wait()

	if err != nil {
		return err
	}

	return nil
}

func (s *ArticleService) DeleteArticle(id string) error {
	wg := sync.WaitGroup{}
	var err error
	wg.Add(2)

	go func() {
		defer wg.Done()
		_, err = s.ArticleRepository.FindArticleByID(id)
		if err != nil {
			err = errors.New("article not found")
		}
	}()

	go func() {
		defer wg.Done()
		err = s.ArticleRepository.DeleteArticle(id)
		if err != nil {
			err = errors.New("failed to delete article")
		}
	}()

	wg.Wait()

	if err != nil {
		return err
	}

	return nil
}
