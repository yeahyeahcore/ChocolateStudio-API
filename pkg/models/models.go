package models

import (
	"time"

	"gorm.io/gorm"
)

type Photo_image struct {
	ID        uint `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	PhotoID   uint           `json:"photo_id,omitempty"`
	ImageURL  string         `json:"image_url,omitempty"`
}

type Book_image struct {
	ID        uint `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	BookID    uint           `json:"book_id,omitempty"`
	ImageURL  string         `json:"image_url,omitempty"`
}

type Photography struct {
	ID        uint `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Score     int            `json:"score,omitempty"`
	Title     string         `json:"title,omitempty"`
	Text      string         `json:"text,omitempty"`
	Price     uint           `json:"price,omitempty"`
	PriceText string         `json:"price_text,omitempty"`
}

type PhotoBook struct {
	ID        uint `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Score     int            `json:"score,omitempty"`
	Title     string         `json:"title,omitempty"`
	Text      string         `json:"text,omitempty"`
	Price     uint           `json:"price,omitempty"`
	PriceText string         `json:"price_text,omitempty"`
}
