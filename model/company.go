package model

import (
	"time"

	"gorm.io/gorm"
)

// Bảng quản lý thông tin công ty// Company struct represents the company entity
type Company struct {
	ID             uint           `json:"id" gorm:"primaryKey"`                // ID
	Name           string         `json:"name"`                                // Tên công ty
	Introduction   string         `json:"introduction"`                        // Giới thiệu về công ty
	Owner          string         `json:"owner"`                               // Người đại diện
	SIC            string         `json:"sic"`                                 // Mã SIC
	ICB            string         `json:"icb"`                                 // Mã ngành ICB
	Category       string         `json:"category"`                            // Loại hình doanh nghiệp
	Major          string         `json:"major"`                               // Tên ngành hoạt động
	MajorCode      string         `json:"major_code"`                          // Mã ngành
	TaxCode        string         `json:"tax_code"`                            // Mã số thuế
	
	BirthDate      time.Time      `json:"birth_date"`                          // Ngày thành lập
	CharterCapital int64          `json:"charter_capital"`                     // Vốn điều lệ
	Employees      int            `json:"employees"`                           // Số lượng nhân viên
	Branches       int            `json:"branches"`                            // Số lượng chi nhánh
	Address        string         `json:"address"`                             // Địa chỉ
	Activity       string         `json:"activity"`                            // Tình trạng hoạt động
	
	ListingDate    time.Time      `json:"listing_date"`                        // Ngày niêm yết
	ListedFloor    string         `json:"listed_floor"`                        // Nơi niêm yết
	IPOPrice       int64          `json:"ipo_price"`                           // Giá chào sàn
	ListedVolume   int64          `json:"listed_volume"`                       // Khối lượng niêm yết
	MarketCap      int64          `json:"market_cap"`                          // Thị giá vốn
	SLCP           int64          `json:"slcp"`                                // SLCP lưu hành

	CreatedAt      time.Time      `json:"created_at" swaggerignore:"true"`     // Thời gian tạo
	DeletedAt      gorm.DeletedAt `json:"-" swaggerignore:"true" gorm:"index"` // Thời gian xóa (soft delete)
	UpdatedAt      time.Time      `json:"updated_at" swaggerignore:"true"`     // Thời gian cập nhật
}
