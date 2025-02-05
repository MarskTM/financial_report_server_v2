package model

import (
	"time"

	"gorm.io/gorm"
)

type UserReport struct {
	ID         int32 `json:"id" gorm:"primarykey"`
	ProfileID  int32 `json:"profile_id"`
	DocumentID int32 `json:"document_id"`

	Name     string    `json:"name"`     // Tên báo cáo.
	Category string    `json:"category"` // Loại hình báo cáo.
	Date     time.Time `json:"date"`     // Ngày công bố báo cáo.

	Reports  []FinancialReport `json:"reports" gorm:"foreignKey:UserReportID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Document Document          `json:"document"  gorm:"foreignKey:DocumentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time      `json:"created_at" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updated_at" swaggerignore:"true"`
}
