package model

import (
	"time"

	"gorm.io/gorm"
)

// Bảng quản lý thông tin công ty// Company struct represents the company entity
type Company struct {
	ID int32 `json:"id" gorm:"primary_key"`
	// Muc 1: thông tin cơ bản
	CompanyName       string    `json:"company_name"`        // Tên công ty
	TaxCode           string    `json:"tax_code"`            // Mã số thuế
	CompanyCode       string    `json:"company_code"`        // Mã công ty
	EstablishmentDate time.Time `json:"establishment_date" ` // Ngày thành lập

	CompanyType        string `json:"company_type"`        // Loại hình công ty
	CompanyLogo        string `json:"company_logo"`        // Đường dân logo
	CompanyAddress     string `json:"company_address"`     // Địa chỉ công ty
	CompanyEmail       string `json:"company_email"`       // E-mail công ty
	CompanyPhone       string `json:"company_phone"`       // E-phone công ty
	CompanyWebsite     string `json:"company_website"`     // Đường dẫn trang chủ
	CompanyDescription string `json:"company_description"` // Giới thiệu về công ty

	// Mục 3: Thông tin niêm yết
	MarketCapitalization float64   `json:"market_capitalization"`  // Vốn hóa
	FirstTradingDate     time.Time `json:"first_trading_date"`     // Ngày giao dịch đầu tiên
	FirstTradingPrice    float64   `json:"first_trading_price"`    // Giá giao dịch ngày đầu
	InitialListingVolume float64   `json:"initial_listing_volume"` // Khối lượng niêm yết lần đầu
	CurrentListingVolume float64   `json:"current_listing_volume"` // Khối lượng niêm yết hiện tại
	OutstandingShares    float64   `json:"outstanding_shares"`     // Khối lượng cổ phiếu đang lưu hành

	CompanyStakeholders []CompanyStakeholder `json:"company_stakeholder" gorm:"foreignKey:CompanyID"`

	CreatedAt time.Time      `json:"created_at" swaggerignore:"true"`     // Thời gian tạo
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true" gorm:"index"` // Thời gian xóa (soft delete)
	UpdatedAt time.Time      `json:"updated_at" swaggerignore:"true"`     // Thời gian cập nhật
}
