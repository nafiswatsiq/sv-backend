package service

import (
	"context"
	"errors"
	"sharingvision/internal/domain/article"
	"sharingvision/internal/repository"
	"strings"
)

type ArticleService interface {
	Create(ctx context.Context, input article.CreateArticleRequest) (*article.Article, error)
	List(ctx context.Context, limit, offset int) ([]article.Article, error)
	GetByID(ctx context.Context, id uint) (*article.Article, error)
	Update(ctx context.Context, id uint, input article.UpdateArticleRequest) (*article.Article, error)
	Delete(ctx context.Context, id uint) error
}

type articleService struct {
	repository repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{repository: repo}
}

func (s *articleService) Create(ctx context.Context, input article.CreateArticleRequest) (*article.Article, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	post := &article.Article{
		Title:    input.Title,
		Content:  input.Content,
		Category: input.Category,
		Status:   strings.ToLower(input.Status),
	}

	if err := s.repository.Create(ctx, post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *articleService) List(ctx context.Context, limit, offset int) ([]article.Article, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return s.repository.List(ctx, limit, offset)
}

func (s *articleService) GetByID(ctx context.Context, id uint) (*article.Article, error) {
	if id == 0 {
		return nil, errors.New("article id is required")
	}
	return s.repository.GetByID(ctx, id)
}

func (s *articleService) Update(ctx context.Context, id uint, input article.UpdateArticleRequest) (*article.Article, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	existing, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	existing.Title = input.Title
	existing.Content = input.Content
	existing.Category = input.Category
	existing.Status = strings.ToLower(input.Status)

	if err := s.repository.Update(ctx, existing); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *articleService) Delete(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("article id is required")
	}
	return s.repository.Delete(ctx, id)
}
