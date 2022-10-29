package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Like struct {
	ID        string `json:"id" gorm:"primaryKey"`
	ArticleID string `json:"article_id" gorm:"not null"`
	UserID    string `json:"user_id" gorm:"not null"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt int64  `json:"deleted_at" gorm:"index"`
}

type Dislike struct {
	ID        string `json:"id" gorm:"primaryKey"`
	ArticleID string `json:"article_id" gorm:"not null"`
	UserID    string `json:"user_id" gorm:"not null"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt int64  `json:"deleted_at" gorm:"index"`
}

func (l *Like) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = uuid.New().String()
	return
}

func (d *Dislike) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.New().String()
	return
}
