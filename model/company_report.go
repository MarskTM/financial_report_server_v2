package model

import (
	"time"

	"gorm.io/gorm"
)

type CompanyReport struct {
	ID         int32 `json:"id" gorm:"primarykey"`
	CompanyID  int32 `json:"company_id"`
	DocumentID int32 `json:"document_id"`

	Name     string    `json:"name"`     // Tên báo cáo.
	Category string    `json:"category"` // Loại hình báo cáo.
	Date     time.Time `json:"date"`     // Ngày công bố báo cáo.

	Reports   []FinancialReport `json:"reports" gorm:"foreignKey:CompanyReportID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Document Document          `json:"document"  gorm:"foreignKey:DocumentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}
