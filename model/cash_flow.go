package model

import (
	"time"

	"gorm.io/gorm"
)

type CashflowStatement struct {
	ID       int32 `json:"id" gorm:"primarykey"` // ID của báo cáo
	ReportID int32 `json:"report_id"`

	// Lưu chuyển tiền tệ từ hoạt động kinh doanh
	NetCashFlowFromOperatingActivities          *float64 `json:"net_cash_flow_from_operating_activities"`            // Lưu chuyển tiền thuần từ hoạt động kinh doanh
	NetCashFlowFromOperatingActivitiesBeforeTax *float64 `json:"net_cash_flow_from_operating_activities_before_tax"` // Lưu chuyển tiền thuần từ hoạt động kinh doanh trước thuế
	ProfitLossBeforeChangesInWorkingCapital     *float64 `json:"profit_loss_before_changes_in_working_capital"`      // Lợi nhuận trước thay đổi vốn lưu động
	InterestAndSimilarIncome                    *float64 `json:"interest_and_similar_income"`                        // Thu nhập từ lãi và các khoản tương tự
	InterestAndSimilarExpenses                  *float64 `json:"interest_and_similar_expenses"`                      // Chi phí lãi và các khoản tương tự
	FeeAndCommissionIncomeReceived              *float64 `json:"fee_and_commission_income_received"`                 // Thu nhập từ phí và hoa hồng đã nhận
	IncomeFromTradingSecurities                 *float64 `json:"income_from_trading_securities"`                     // Thu nhập từ giao dịch chứng khoán
	OtherIncome                                 *float64 `json:"other_income"`                                       // Thu nhập khác
	RecoveredBadDebts                           *float64 `json:"recovered_bad_debts"`                                // Thu hồi nợ xấu
	PaymentsToEmployeesAndSuppliers             *float64 `json:"payments_to_employees_and_suppliers"`                // Thanh toán cho nhân viên và nhà cung cấp
	CorporateIncomeTaxPaid                      *float64 `json:"corporate_income_tax_paid"`                          // Thuế thu nhập doanh nghiệp đã nộp

	// Biến động tài sản và nợ ngắn hạn
	DepositsAtStateBank                                                *float64 `json:"deposits_at_state_bank"`                                                      // Tiền gửi tại ngân hàng nhà nước
	IncreaseDecreaseInDepositsAndLoansToOtherCreditInstitutions        *float64 `json:"increase_decrease_in_deposits_and_loans_to_other_credit_institutions"`        // Biến động tiền gửi và khoản vay tại các tổ chức tín dụng khác
	IncreaseDecreaseInTradingSecurities                                *float64 `json:"increase_decrease_in_trading_securities"`                                     // Biến động chứng khoán kinh doanh
	IncreaseDecreaseInDerivativesAndOtherFinancialAssets               *float64 `json:"increase_decrease_in_derivatives_and_other_financial_assets"`                 // Biến động công cụ phái sinh và tài sản tài chính khác
	IncreaseDecreaseInLoansToCustomers                                 *float64 `json:"increase_decrease_in_loans_to_customers"`                                     // Biến động các khoản cho vay khách hàng
	IncreaseDecreaseInInterestAndFeeReceivables                        *float64 `json:"increase_decrease_in_interest_and_fee_receivables"`                           // Biến động lãi và phí phải thu
	IncreaseDecreaseInProvisionFunds                                   *float64 `json:"increase_decrease_in_provision_funds"`                                        // Biến động các quỹ dự phòng
	IncreaseDecreaseInOtherOperatingAssets                             *float64 `json:"increase_decrease_in_other_operating_assets"`                                 // Biến động các tài sản hoạt động khác
	IncreaseDecreaseInGovernmentAndStateBankBorrowings                 *float64 `json:"increase_decrease_in_government_and_state_bank_borrowings"`                   // Biến động các khoản vay từ chính phủ và ngân hàng nhà nước
	IncreaseDecreaseInDepositsAndBorrowingsFromOtherCreditInstitutions *float64 `json:"increase_decrease_in_deposits_and_borrowings_from_other_credit_institutions"` // Biến động tiền gửi và khoản vay từ các tổ chức tín dụng khác
	IncreaseDecreaseInCustomerDeposits                                 *float64 `json:"increase_decrease_in_customer_deposits"`                                      // Biến động tiền gửi của khách hàng
	IncreaseDecreaseInDerivativesAndOtherFinancialLiabilities          *float64 `json:"increase_decrease_in_derivatives_and_other_financial_liabilities"`            // Biến động công cụ phái sinh và nợ tài chính khác
	IncreaseDecreaseInTrustFundAndInvestmentsFromGovernmentAndOthers   *float64 `json:"increase_decrease_in_trust_fund_and_investments_from_government_and_others"`  // Biến động quỹ tín thác và đầu tư từ chính phủ và các bên khác
	IncreaseDecreaseInIssuedDebtSecurities                             *float64 `json:"increase_decrease_in_issued_debt_securities"`                                 // Biến động công cụ nợ đã phát hành
	IncreaseDecreaseInInterestAndFeePayables                           *float64 `json:"increase_decrease_in_interest_and_fee_payables"`                              // Biến động lãi và phí phải trả
	IncreaseDecreaseInOtherOperatingLiabilities                        *float64 `json:"increase_decrease_in_other_operating_liabilities"`                            // Biến động các khoản nợ hoạt động khác

	// Lưu chuyển tiền tệ từ hoạt động đầu tư
	NetCashFlowFromInvestingActivities                     *float64 `json:"net_cash_flow_from_investing_activities"`                          // Lưu chuyển tiền thuần từ hoạt động đầu tư
	PaymentsForPurchaseOfFixedAssetsAndOtherLongTermAssets *float64 `json:"payments_for_purchase_of_fixed_assets_and_other_long_term_assets"` // Chi trả để mua tài sản cố định và tài sản dài hạn khác
	ProceedsFromDisposalOfFixedAssets                      *float64 `json:"proceeds_from_disposal_of_fixed_assets"`                           // Thu nhập từ việc bán tài sản cố định
	PaymentsForDisposalOfFixedAssets                       *float64 `json:"payments_for_disposal_of_fixed_assets"`                            // Chi trả cho việc thanh lý tài sản cố định
	InvestmentsInOtherEntities                             *float64 `json:"investments_in_other_entities"`                                    // Đầu tư vào các tổ chức khác
	ProceedsFromDisposalOfInvestmentsInOtherEntities       *float64 `json:"proceeds_from_disposal_of_investments_in_other_entities"`          // Thu nhập từ việc thanh lý đầu tư vào tổ chức khác
	DividendsAndInterestReceived                           *float64 `json:"dividends_and_interest_received"`                                  // Cổ tức và thu nhập lãi đã nhận

	// Lưu chuyển tiền tệ từ hoạt động tài chính
	NetCashFlowFromFinancingActivities                    *float64 `json:"net_cash_flow_from_financing_activities"`                        // Lưu chuyển tiền thuần từ hoạt động tài chính
	ProceedsFromIssuanceOfSharesAndCapitalContributions   *float64 `json:"proceeds_from_issuance_of_shares_and_capital_contributions"`     // Thu nhập từ phát hành cổ phiếu và góp vốn
	ProceedsFromIssuanceOfLongTermDebtSecuritiesAndLoans  *float64 `json:"proceeds_from_issuance_of_long_term_debt_securities_and_loans"`  // Thu nhập từ phát hành công cụ nợ dài hạn và vay
	PaymentsForSettlementOfLongTermDebtSecuritiesAndLoans *float64 `json:"payments_for_settlement_of_long_term_debt_securities_and_loans"` // Chi trả nợ dài hạn và công cụ nợ
	DividendsPaid                                         *float64 `json:"dividends_paid"`                                                 // Cổ tức đã chi trả

	// Tổng hợp lưu chuyển tiền tệ
	NetIncreaseDecreaseInCashAndCashEquivalents *float64 `json:"net_increase_decrease_in_cash_and_cash_equivalents"` // Biến động thuần tiền và tương đương tiền
	CashAndCashEquivalentsAtBeginningOfPeriod   *float64 `json:"cash_and_cash_equivalents_at_beginning_of_period"`   // Tiền và tương đương tiền đầu kỳ
	EffectOfExchangeRateFluctuations            *float64 `json:"effect_of_exchange_rate_fluctuations"`               // Ảnh hưởng của biến động tỷ giá hối đoái
	CashAndCashEquivalentsAtEndOfPeriod         *float64 `json:"cash_and_cash_equivalents_at_end_of_period"`         // Tiền và tương đương tiền cuối kỳ

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
