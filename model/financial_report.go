package model

import (
	"time"

	"gorm.io/gorm"
)

type FinancialReport struct {
	// ===================================================== Meta data ============================================================
	ID         int32 `json:"id" gorm:"primarykey"`
	CompanyID  int32 `json:"company_id"`
	DocumentID int32 `json:"document_id"`

	Name     string    `json:"name"`     // Tên báo cáo.
	Category string    `json:"category"` // Loại hình báo cáo.
	Date     time.Time `json:"date"`     // Ngày công bố báo cáo.
	Quarter  string    `json:"quarter"`  // Quý báo cáo

	// ======================================================= Income statement ==========================================================================
	// Thu nhập và chi phí từ lãi
	// InterestAndSimilarIncome   *float64 `json:"interest_and_similar_income"`   // Thu nhập từ lãi và các khoản tương tự
	// InterestAndSimilarExpenses *float64 `json:"interest_and_similar_expenses"` // Chi phí lãi và các khoản tương tự
	// NetInterestIncome          *float64 `json:"net_interest_income"`           // Thu nhập lãi thuần

	OperatingInterestIncome    *float64 `json:"operating_interest_income"`     // Thu nhập lãi từ hoạt động kinh doanh
	OperatingInterestExpenses  *float64 `json:"operating_interest_expenses"`   // Chi phí lãi từ hoạt động kinh doanh
	NetOperatingInterestIncome *float64 `json:"net_operating_interest_income"` // Thu nhập lãi thuần từ hoạt động kinh doanh

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

	// ================================================================= Cashflow Statement =================================================================
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
	IncreaseDecreaseInDepositsAndLoansToOtherCreditInstitutions        *float64 `json:"increase_decrease_in_deposits_and_loans_to_other_credit"`                     // Biến động tiền gửi và khoản vay tại các tổ chức tín dụng khác
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

	// ================================================================= Balance Sheet =================================================================
	// Tổng tài sản
	// TotalAssets *float64 `json:"total_assets"` // Tổng tài sản
	TotalCompanyAssets *float64 `json:"total_company_assets"` // Tổng tài sản

	// Tiền mặt và tương đương tiền
	CashGoldSilver *float64 `json:"cash_gold_silver"` // Tiền mặt, vàng, bạc
	// DepositsAtStateBank          *float64 `json:"deposits_at_state_bank"`          // Tiền gửi tại ngân hàng nhà nước
	// DepositsAtCreditInstitutions *float64 `json:"deposits_at_credit_institutions"` // Tiền gửi tại các tổ chức tín dụng
	StateBankDeposits         *float64 `json:"state_bank_deposits"`         // Tiền gửi tại ngân hàng nhà nước
	CreditInstitutionDeposits *float64 `json:"credit_institution_deposits"` // Tiền gửi tại các tổ chức tín dụng

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
	// TotalLiabilities                        *float64 `json:"total_liabilities"`                           // Tổng nợ phải trả
	CompanyLiabilities *float64 `json:"company_liabilities"` // Tổng nợ

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

	// ================================================================= Financial Analysis =================================================================
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

	// IncomeStatement   IncomeStatement   `json:"income_statement" gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// CashFlowStatement CashflowStatement `json:"cash_flow" gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// BalanceSheet      BalanceSheet      `json:"balance_sheet" gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// FinancialAnalyst  FinancialAnalyzer `json:"financial_analyst" gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
