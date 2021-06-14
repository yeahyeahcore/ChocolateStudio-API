package models

type Photo_image struct {
	ID        uint `gorm:"primarykey" json:"id,omitempty"`
	PhotoID   uint           `json:"photo_id,omitempty"`
	ImageURL  string         `json:"image_url,omitempty"`
}

type Book_image struct {
	ID        uint `gorm:"primarykey" json:"id,omitempty"`
	BookID    uint           `json:"book_id,omitempty"`
	ImageURL  string         `json:"image_url,omitempty"`
}

type Photography struct {
	ID        uint `gorm:"primarykey" json:"id,omitempty"`
	Score     int            `json:"score,omitempty"`
	Title     string         `json:"title,omitempty"`
	Text      string         `json:"text,omitempty"`
	Price     float32           `json:"price,omitempty"`
	PriceText string         `json:"price_text,omitempty"`
}

type PhotoBook struct {
	ID        uint `gorm:"primarykey" json:"id,omitempty"`
	Score     int            `json:"score,omitempty"`
	Title     string         `json:"title,omitempty"`
	Text      string         `json:"text,omitempty"`
	Price     float32           `json:"price,omitempty"`
	PriceText string         `json:"price_text,omitempty"`
}
