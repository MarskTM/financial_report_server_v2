package model

import (
	"time"

	"gorm.io/gorm"
)

// Bảng quản lý thông tin công ty
type Company struct {
	ID        int32  `json:"id"`
	CompanyID int32  `json:"company_id"`
	Owner     string `json:"owner"` // Tên cổ đông
	Type      string `json:"type"`  // Cá nhân/Tổ chức

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}
