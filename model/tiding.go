package model

import (
	"time"

	"gorm.io/gorm"
)

type Tiding struct {
	ID       int32    `json:"id"`
	ParentID int32    `json:"parent_id"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	Images   []string `json:"images"`

	PrevContent string `json:"prev_content"`
	PrevImage   string `json:"prev_image"`

	SubTidings []Tiding `json:"sub_tidings" gorm:"foreignKey:ParentID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	CreatedAt time.Time      `json:"created_at" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updated_at" swaggerignore:"true"`
}
