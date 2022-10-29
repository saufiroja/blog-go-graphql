package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Body      string `json:"body" gorm:"not null"`
	ArticleID string `json:"article_id" gorm:"not null"`
	UserID    string `json:"user_id" gorm:"not null"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt int64  `json:"deleted_at" gorm:"index"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
