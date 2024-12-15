package model

import (
	"time"

	"gorm.io/gorm"
)

type FinancialAnalyzer struct {
	ID       int32 `json:"id" gorm:"primarykey"`
	ReportID int32 `json:"report_id"`

	// Nhóm chỉ số Định giá
	NetIncome         float64  `json:"net_income"`          // Lợi nhuận sau thuế
	TotalAssets       float64  `json:"total_assets"`        // Tổng tài sản
	TotalLiabilities  float64  `json:"total_liabilities"`   // Tổng nợ
	OutstandingShares float64  `json:"outstanding_shares"`  // Số lượng cổ phiếu lưu hành
	StockPrice        float64  `json:"stock_price"`         // Giá cổ phiếu
	DividendsPerShare *float64 `json:"dividends_per_share"` // Cổ tức trên mỗi cổ phiếu
	EPS               *float64 `json:"eps"`                 // Thu nhập trên mỗi cổ phiếu (EPS)
	BVPS              *float64 `json:"bvps"`                // Giá trị sổ sách trên mỗi cổ phiếu (BVPS)
	PE                *float64 `json:"pe"`                  // Tỷ lệ giá trên thu nhập (P/E)
	PB                *float64 `json:"pb"`                  // Tỷ lệ giá trên giá trị sổ sách (P/B)
	DividendYield     *float64 `json:"dividend_yield"`      // Tỷ suất cổ tức

	// Nhóm chỉ số Tỷ suất
	ROE  *float64 `json:"roe"`  // Tỷ suất lợi nhuận trên vốn chủ sở hữu (ROE)
	ROA  *float64 `json:"roa"`  // Tỷ suất sinh lời trên tổng tài sản (ROA)
	YOEA *float64 `json:"yoea"` // Tỷ suất sinh lợi của Tài sản Có sinh lãi (YOEA)
	COF  *float64 `json:"cof"`  // Tỷ lệ chi phí hình thành Tài sản Có sinh lãi (COF)
	NIM  *float64 `json:"nim"`  // Tỷ lệ thu nhập lãi thuần (NIM)

	// Nhóm chỉ số Tăng trưởng
	PreTaxProfitGrowth  *float64 `json:"pre_tax_profit_growth"`  // Tăng trưởng lợi nhuận trước thuế
	PostTaxProfitGrowth *float64 `json:"post_tax_profit_growth"` // Tăng trưởng lợi nhuận sau thuế
	TotalAssetsGrowth   *float64 `json:"total_assets_growth"`    // Tăng trưởng tổng tài sản
	EquityGrowth        *float64 `json:"equity_growth"`          // Tăng trưởng vốn chủ sở hữu
	LoanGrowth          *float64 `json:"loan_growth"`            // Tăng trưởng dư nợ cho vay

	// Nhóm chỉ số Thanh khoản và An toàn
	LDR                  *float64 `json:"ldr"`                     // Dư nợ cho vay khách hàng/Tổng vốn huy động (LDR)
	LoanToTotalAssets    *float64 `json:"loan_to_total_assets"`    // Dư nợ cho vay/Tổng tài sản Có
	EquityToDepositRatio *float64 `json:"equity_to_deposit_ratio"` // Vốn chủ sở hữu/Tổng vốn huy động
	EquityToTotalAssets  *float64 `json:"equity_to_total_assets"`  // Vốn chủ sở hữu/Tổng tài sản Có

	// Nhóm chỉ số Rủi ro
	RiskReserveToLoans         *float64 `json:"risk_reserve_to_loans"`          // Dự phòng rủi ro tín dụng/Tổng dư nợ
	EarningAssetsToTotalAssets *float64 `json:"earning_assets_to_total_assets"` // Tài sản Có sinh lãi/Tổng tài sản Có

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
