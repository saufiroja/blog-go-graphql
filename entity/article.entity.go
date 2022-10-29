package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	PhotoURL     string    `json:"photo_url" gorm:"not null"`
	Title        string    `json:"title" gorm:"not null"`
	Body         string    `json:"body" gorm:"not null"`
	Category     string    `json:"category" gorm:"not null"`
	UserID       string    `json:"user_id" gorm:"not null"`
	TotalComment int       `json:"total_comment" gorm:"-"`
	TotalLike    int       `json:"total_like" gorm:"-"`
	TotalDislike int       `json:"total_dislike" gorm:"-"`
	Comments     []Comment `json:"comments" gorm:"foreignKey:ArticleID"`
	Likes        []Like    `json:"likes" gorm:"foreignKey:ArticleID"`
	Dislikes     []Dislike `json:"dislikes" gorm:"foreignKey:ArticleID"`
	CreatedAt    int64     `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt    int64     `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt    int64     `json:"deleted_at" gorm:"index"`
}

func (a *Article) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New().String()
	return
}
