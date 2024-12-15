package model

import (
	"time"

	"gorm.io/gorm"
)

type BalanceSheet struct {
	ID       int32 `json:"id" gorm:"primarykey"`
	ReportID int32 `json:"report_id"`
	
	// Tổng tài sản
	TotalAssets *float64 `json:"total_assets"` // Tổng tài sản

	// Tiền mặt và tương đương tiền
	CashGoldSilver               *float64 `json:"cash_gold_silver"`                // Tiền mặt, vàng, bạc
	DepositsAtStateBank          *float64 `json:"deposits_at_state_bank"`          // Tiền gửi tại ngân hàng nhà nước
	DepositsAtCreditInstitutions *float64 `json:"deposits_at_credit_institutions"` // Tiền gửi tại các tổ chức tín dụng

	// Chứng khoán và tài sản đầu tư
	TradingSecurities                  *float64 `json:"trading_securities"`                     // Chứng khoán kinh doanh
	ProvisionTradingSecurities         *float64 `json:"provision_trading_securities"`           // Dự phòng chứng khoán kinh doanh
	DerivativesAndOtherFinancialAssets *float64 `json:"derivatives_and_other_financial_assets"` // Các công cụ phái sinh và tài sản tài chính khác
	InvestmentSecurities               *float64 `json:"investment_securities"`                  // Chứng khoán đầu tư
	SecuritiesAvailableForSale         *float64 `json:"securities_available_for_sale"`          // Chứng khoán sẵn sàng để bán
	SecuritiesHeldToMaturity           *float64 `json:"securities_held_to_maturity"`            // Chứng khoán nắm giữ đến ngày đáo hạn
	ProvisionInvestmentSecurities      *float64 `json:"provision_investment_securities"`        // Dự phòng chứng khoán đầu tư

	// Các khoản đầu tư dài hạn
	LongTermInvestments          *float64 `json:"long_term_investments"`           // Đầu tư dài hạn
	InvestmentsInSubsidiaries    *float64 `json:"investments_in_subsidiaries"`     // Đầu tư vào công ty con
	InvestmentsInJointVentures   *float64 `json:"investments_in_joint_ventures"`   // Đầu tư vào công ty liên kết, liên doanh
	OtherLongTermInvestments     *float64 `json:"other_long_term_investments"`     // Đầu tư dài hạn khác
	ProvisionLongTermInvestments *float64 `json:"provision_long_term_investments"` // Dự phòng đầu tư dài hạn

	// Tài sản cố định
	FixedAssets           *float64 `json:"fixed_assets"`            // Tài sản cố định
	TangibleFixedAssets   *float64 `json:"tangible_fixed_assets"`   // Tài sản cố định hữu hình
	LeasedFixedAssets     *float64 `json:"leased_fixed_assets"`     // Tài sản cố định thuê tài chính
	IntangibleFixedAssets *float64 `json:"intangible_fixed_assets"` // Tài sản cố định vô hình
	InvestmentProperties  *float64 `json:"investment_properties"`   // Bất động sản đầu tư

	// Tài sản khác
	OtherAssets *float64 `json:"other_assets"` // Tài sản khác

	// Tổng nguồn vốn
	TotalLiabilitiesAndEquity *float64 `json:"total_liabilities_and_equity"` // Tổng nợ phải trả và vốn chủ sở hữu

	// Nợ phải trả
	TotalLiabilities                        *float64 `json:"total_liabilities"`                           // Tổng nợ phải trả
	GovernmentDebtsAndStateBank             *float64 `json:"government_debts_and_state_bank"`             // Nợ chính phủ và ngân hàng nhà nước
	DepositsAndLoansFromCreditInstitutions  *float64 `json:"deposits_and_loans_from_credit_institutions"` // Tiền gửi và vay từ các tổ chức tín dụng
	CustomerDeposits                        *float64 `json:"customer_deposits"`                           // Tiền gửi của khách hàng
	DerivativesAndOtherFinancialLiabilities *float64 `json:"derivatives_and_other_financial_liabilities"` // Các công cụ phái sinh và nợ tài chính khác
	SponsoredCapitalAndTrustInvestments     *float64 `json:"sponsored_capital_and_trust_investments"`     // Vốn tài trợ và ủy thác đầu tư
	IssuedDebtInstruments                   *float64 `json:"issued_debt_instruments"`                     // Công cụ nợ đã phát hành
	OtherLiabilities                        *float64 `json:"other_liabilities"`                           // Nợ phải trả khác

	// Vốn chủ sở hữu
	Equity                        *float64 `json:"equity"`                           // Vốn chủ sở hữu
	CreditInstitutionCapital      *float64 `json:"credit_institution_capital"`       // Vốn của tổ chức tín dụng
	CharterCapital                *float64 `json:"charter_capital"`                  // Vốn điều lệ
	CapitalForConstruction        *float64 `json:"capital_for_construction"`         // Vốn đầu tư xây dựng
	SharePremium                  *float64 `json:"share_premium"`                    // Thặng dư vốn cổ phần
	TreasuryShares                *float64 `json:"treasury_shares"`                  // Cổ phiếu quỹ
	PreferredShares               *float64 `json:"preferred_shares"`                 // Cổ phiếu ưu đãi
	OtherEquity                   *float64 `json:"other_equity"`                     // Vốn khác
	FundsOfCreditInstitution      *float64 `json:"funds_of_credit_institution"`      // Các quỹ của tổ chức tín dụng
	ForeignExchangeDifference     *float64 `json:"foreign_exchange_difference"`      // Chênh lệch tỷ giá hối đoái
	RevaluationDifferenceOfAssets *float64 `json:"revaluation_difference_of_assets"` // Chênh lệch đánh giá lại tài sản
	UndistributedProfits          *float64 `json:"undistributed_profits"`            // Lợi nhuận chưa phân phối

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}
