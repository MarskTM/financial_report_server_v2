package model

import (
	"time"

	"gorm.io/gorm"
)

type Document struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	ShortName   string `json:"short_name"`
	Cdn         string `json:"cdn"`
	Description string `json:"description"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `json:"-"`
}


