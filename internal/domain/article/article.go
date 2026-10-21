package article

import (
	"errors"
	"strings"
	"time"
)

type Status string

const (
	StatusPublish Status = "publish"
	StatusDraft   Status = "draft"
	StatusTrash   Status = "thrash"
)

type Article struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(200);not null" json:"title"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	Category    string    `gorm:"type:varchar(100);not null" json:"category"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"updated_date"`
	Status      string    `gorm:"type:varchar(100);not null" json:"status"`
}

func (Article) TableName() string {
	return "posts"
}

type CreateArticleRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

type UpdateArticleRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

func (r CreateArticleRequest) Validate() error {
	trimmedTitle := strings.TrimSpace(r.Title)
	trimmedContent := strings.TrimSpace(r.Content)
	trimmedCategory := strings.TrimSpace(r.Category)
	trimmedStatus := strings.TrimSpace(strings.ToLower(r.Status))

	if len(trimmedTitle) < 20 {
		return errors.New("title is required and minimum 20 characters")
	}
	if len(trimmedContent) < 200 {
		return errors.New("content is required and minimum 200 characters")
	}
	if len(trimmedCategory) < 3 {
		return errors.New("category is required and minimum 3 characters")
	}
	if trimmedStatus == "" {
		return errors.New("status is required")
	}
	if !isValidStatus(trimmedStatus) {
		return errors.New("status must be one of publish, draft, thrash")
	}

	r.Title = trimmedTitle
	r.Content = trimmedContent
	r.Category = trimmedCategory
	r.Status = trimmedStatus
	return nil
}

func (r UpdateArticleRequest) Validate() error {
	trimmedTitle := strings.TrimSpace(r.Title)
	trimmedContent := strings.TrimSpace(r.Content)
	trimmedCategory := strings.TrimSpace(r.Category)
	trimmedStatus := strings.TrimSpace(strings.ToLower(r.Status))

	if len(trimmedTitle) < 20 {
		return errors.New("title is required and minimum 20 characters")
	}
	if len(trimmedContent) < 200 {
		return errors.New("content is required and minimum 200 characters")
	}
	if len(trimmedCategory) < 3 {
		return errors.New("category is required and minimum 3 characters")
	}
	if trimmedStatus == "" {
		return errors.New("status is required")
	}
	if !isValidStatus(trimmedStatus) {
		return errors.New("status must be one of publish, draft, thrash")
	}

	r.Title = trimmedTitle
	r.Content = trimmedContent
	r.Category = trimmedCategory
	r.Status = trimmedStatus
	return nil
}

func isValidStatus(status string) bool {
	switch Status(strings.ToLower(status)) {
	case StatusPublish, StatusDraft, StatusTrash:
		return true
	default:
		return false
	}
}
