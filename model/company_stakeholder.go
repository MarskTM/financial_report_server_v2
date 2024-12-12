package model

import (
	"time"

	"gorm.io/gorm"
)

type CompanyStakeholder struct {
	ID        int32     `json:"id"`
	CompanyID int32     `json:"company_id"`
	Avatar    string    `json:"avatar"`
	Name      string    `json:"name"`
	Position  string    `json:"position"`
	YearStart time.Time `json:"year_start"` // ngày bổ nhiệm

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}
