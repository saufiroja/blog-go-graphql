package dto

type CreateArticle struct {
	PhotoURL string `json:"photo_url" gorm:"not null"`
	Title    string `json:"title" gorm:"not null"`
	Body     string `json:"body" gorm:"not null"`
	Category string `json:"category" gorm:"not null"`
	UserID   string `json:"user_id" gorm:"not null"`
}

type UpdateArticle struct {
	PhotoURL string `json:"photo_url" gorm:"not null"`
	Title    string `json:"title" gorm:"not null"`
	Body     string `json:"body" gorm:"not null"`
	Category string `json:"category" gorm:"not null"`
	UserID   string `json:"user_id" gorm:"not null"`
}
