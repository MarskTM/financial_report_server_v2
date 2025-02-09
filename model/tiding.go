package model

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Tiding struct {
	ID       int32          `json:"id"`
	ParentID *int32         `json:"parent_id"` // Make ParentID a pointer to int32
	Title    string         `json:"title"`
	Content  string         `json:"content"`
	Category string         `json:"category"`
	Images   pq.StringArray `json:"images" gorm:"type:text[]"`
	State    bool           `json:"state"`

	PrevContent string `json:"prev_content"`
	PrevImage   string `json:"prev_image"`

	SubTidings []Tiding `json:"tidings" gorm:"foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"created_at" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updated_at" swaggerignore:"true"`
}
