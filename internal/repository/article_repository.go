package repository

import (
	"context"
	"sharingvision/internal/domain/article"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(ctx context.Context, articleModel *article.Article) error
	List(ctx context.Context, limit, offset int) ([]article.Article, error)
	GetByID(ctx context.Context, id uint) (*article.Article, error)
	Update(ctx context.Context, articleModel *article.Article) error
	Delete(ctx context.Context, id uint) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) Create(ctx context.Context, articleModel *article.Article) error {
	return r.db.WithContext(ctx).Create(articleModel).Error
}

func (r *articleRepository) List(ctx context.Context, limit, offset int) ([]article.Article, error) {
	var posts []article.Article
	query := r.db.WithContext(ctx).Order("created_date DESC")
	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}
	result := query.Find(&posts)
	return posts, result.Error
}

func (r *articleRepository) GetByID(ctx context.Context, id uint) (*article.Article, error) {
	var post article.Article
	result := r.db.WithContext(ctx).First(&post, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}

func (r *articleRepository) Update(ctx context.Context, articleModel *article.Article) error {
	return r.db.WithContext(ctx).Model(&article.Article{}).Where("id = ?", articleModel.ID).Updates(article.Article{
		Title:    articleModel.Title,
		Content:  articleModel.Content,
		Category: articleModel.Category,
		Status:   articleModel.Status,
	}).Error
}

func (r *articleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&article.Article{}, id).Error
}
