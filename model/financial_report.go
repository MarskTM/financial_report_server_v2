package model

import (
	"time"

	"gorm.io/gorm"
)

type FinancialReport struct {
	ID         string `json:"id"`
	CompanyID  string `json:"company_id"`
	DocumentID string `json:"document_id"`

	Name     string    `json:"name"`     // Tên báo cáo.
	Category string    `json:"category"` // Loại hình báo cáo.
	Quater   int64     `json:"quater"`   // Quý báo cáo.
	Date     time.Time `json:"date"`     // Ngày công bố báo cáo.

	BalenceSheetID    int32 `json:"balance_sheet_id"`
	IncomeStatementID int32 `json:"income_statement_id"`
	CashFlowID        int32 `json:"cash_flow_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
