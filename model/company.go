package model

import (
	"time"

	"gorm.io/gorm"
)

// Bảng quản lý thông tin công ty// Company struct represents the company entity
type Company struct {
	ID int32 `json:"id" gorm:"primary_key"`

	CompanyName     string    `json:"company_name"`      // Tên công ty
	CompanyCode     string    `json:"company_code"`      // Mã công ty
	EstablishedDate time.Time `json:"established_date" ` // Ngày thành lập

	CompanyType    string `json:"company_type"`    // Loại hình công ty
	CompanyLogo    string `json:"company_logo"`    // Đường dân logo
	CompanyAddress string `json:"company_address"` // Địa chỉ công ty

	CompanyStakeholder []CompanyStakeholder `json:"company_stackholder" gorm:"foreignKey:CompanyID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	CreatedAt      time.Time      `json:"created_at" swaggerignore:"true"`     // Thời gian tạo
	DeletedAt      gorm.DeletedAt `json:"-" swaggerignore:"true" gorm:"index"` // Thời gian xóa (soft delete)
	UpdatedAt      time.Time      `json:"updated_at" swaggerignore:"true"`     // Thời gian cập nhật
}
