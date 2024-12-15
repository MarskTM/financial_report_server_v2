package model

import (
	"time"

	"gorm.io/gorm"
)

type IncomeStatement struct {
	ID       int32 `json:"id" gorm:"primarykey"` // ID của báo cáo
	ReportID int32 `json:"report_id"`

	// Thu nhập và chi phí từ lãi
	InterestAndSimilarIncome   *float64 `json:"interest_and_similar_income"`   // Thu nhập từ lãi và các khoản tương tự
	InterestAndSimilarExpenses *float64 `json:"interest_and_similar_expenses"` // Chi phí lãi và các khoản tương tự
	NetInterestIncome          *float64 `json:"net_interest_income"`           // Thu nhập lãi thuần

	// Thu nhập và chi phí từ phí và hoa hồng
	FeeAndCommissionIncome    *float64 `json:"fee_and_commission_income"`     // Thu nhập từ phí và hoa hồng
	FeeAndCommissionExpenses  *float64 `json:"fee_and_commission_expenses"`   // Chi phí phí và hoa hồng
	NetFeeAndCommissionIncome *float64 `json:"net_fee_and_commission_income"` // Thu nhập phí và hoa hồng thuần

	// Thu nhập/lỗ từ giao dịch
	NetGainLossFromForexAndGoldTrading  *float64 `json:"net_gain_loss_from_forex_and_gold_trading"` // Lãi/lỗ ròng từ giao dịch ngoại hối và vàng
	NetGainLossFromTradingSecurities    *float64 `json:"net_gain_loss_from_trading_securities"`     // Lãi/lỗ ròng từ chứng khoán kinh doanh
	NetGainLossFromInvestmentSecurities *float64 `json:"net_gain_loss_from_investment_securities"`  // Lãi/lỗ ròng từ chứng khoán đầu tư

	// Thu nhập và chi phí hoạt động khác
	OtherOperatingIncome             *float64 `json:"other_operating_income"`               // Thu nhập hoạt động khác
	OtherOperatingExpenses           *float64 `json:"other_operating_expenses"`             // Chi phí hoạt động khác
	NetOtherOperatingIncomeExpenses  *float64 `json:"net_other_operating_income_expenses"`  // Thu nhập/chi phí hoạt động khác ròng
	IncomeFromInvestmentInAssociates *float64 `json:"income_from_investment_in_associates"` // Thu nhập từ đầu tư vào công ty liên kết

	// Thu nhập hoạt động tổng cộng và chi phí hoạt động
	TotalOperatingIncome *float64 `json:"total_operating_income"` // Thu nhập hoạt động tổng cộng
	OperatingExpenses    *float64 `json:"operating_expenses"`     // Chi phí hoạt động

	// Lợi nhuận trước dự phòng và thuế
	ProfitBeforeProvisionForCreditLosses *float64 `json:"profit_before_provision_for_credit_losses"` // Lợi nhuận trước dự phòng rủi ro tín dụng
	ProvisionExpensesForCreditLosses     *float64 `json:"provision_expenses_for_credit_losses"`      // Chi phí dự phòng rủi ro tín dụng
	ProfitBeforeTax                      *float64 `json:"profit_before_tax"`                         // Lợi nhuận trước thuế

	// Chi phí thuế thu nhập doanh nghiệp
	CurrentCorporateIncomeTaxExpense  *float64 `json:"current_corporate_income_tax_expense"`  // Chi phí thuế thu nhập hiện tại
	DeferredCorporateIncomeTaxExpense *float64 `json:"deferred_corporate_income_tax_expense"` // Chi phí thuế thu nhập hoãn lại
	CorporateIncomeTaxExpense         *float64 `json:"corporate_income_tax_expense"`          // Tổng chi phí thuế thu nhập doanh nghiệp

	// Lợi nhuận sau thuế
	ProfitAfterTax            *float64 `json:"profit_after_tax"`            // Lợi nhuận sau thuế
	MinorityInterest          *float64 `json:"minority_interest"`           // Lợi ích của cổ đông thiểu số
	ParentCompanyShareholders *float64 `json:"parent_company_shareholders"` // Lợi ích của cổ đông công ty mẹ

	// Các chỉ số tài chính
	BasicEarningsPerShare *float64 `json:"basic_earnings_per_share"` // Lợi nhuận cơ bản trên mỗi cổ phiếu
	EBIT                  *float64 `json:"ebit"`                     // Lợi nhuận trước lãi và thuế
	EBITDA                *float64 `json:"ebitda"`                   // Lợi nhuận trước lãi, thuế, khấu hao và khấu trừ

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
