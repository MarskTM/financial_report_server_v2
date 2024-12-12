package model

import (
	"time"

	"gorm.io/gorm"
)

type CompanyStakeholder struct {
	ID        int64  `json:"id"`
	CompanyID int64  `json:"company_id"`
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
	Position  string `json:"position"`
	YearStart int64  `json:"year_start"` // ngày bổ nhiệm

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}
