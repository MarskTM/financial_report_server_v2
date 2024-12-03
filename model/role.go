package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          int32   `json:"id" gorm:"primaryKey"`
	Code        string  `json:"code" gorm:"unique"`
	Type        string  `json:"type"`
	Description *string `json:"description"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}
