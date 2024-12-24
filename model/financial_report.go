package model

import (
	"time"

	"gorm.io/gorm"
)

type FinancialReport struct {
	// ===================================================== Meta data ============================================================
	ID              int32  `json:"id" gorm:"primarykey"`
	UserReportID    int32  `json:"user_report_id"`
	CompanyReportID int32  `json:"company_report_id"`
	Quarter         string `json:"quarter"` // Quý báo cáo

	// ======================================================= Income statement ==========================================================================
	// Thu nhập và chi phí từ lãi
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
	DepositsAtStateBank             *float64 `json:"deposits_at_state_bank"`             // Tiền gửi tại ngân hàng nhà nước
	ChangeDepositsLoansOthers       *float64 `json:"change_deposits_loans_others"`       // Biến động tiền gửi và khoản vay tại các tổ chức tín dụng khác
	ChangeTradingSecurities         *float64 `json:"change_trading_securities"`          // Biến động chứng khoán kinh doanh
	ChangeDerivativesAssets         *float64 `json:"change_derivatives_assets"`          // Biến động công cụ phái sinh và tài sản tài chính khác
	ChangeCustomerLoans             *float64 `json:"change_customer_loans"`              // Biến động các khoản cho vay khách hàng
	ChangeInterestFeeReceivables    *float64 `json:"change_interest_fee_receivables"`    // Biến động lãi và phí phải thu
	ChangeProvisionFunds            *float64 `json:"change_provision_funds"`             // Biến động các quỹ dự phòng
	ChangeOtherAssets               *float64 `json:"change_other_assets"`                // Biến động các tài sản hoạt động khác
	ChangeGovBorrowings             *float64 `json:"change_gov_borrowings"`              // Biến động các khoản vay chính phủ và ngân hàng nhà nước
	ChangeDepositsLoansInstitutions *float64 `json:"change_deposits_loans_institutions"` // Biến động tiền gửi và vay từ tổ chức tín dụng khác
	ChangeCustomerDeposits          *float64 `json:"change_customer_deposits"`           // Biến động tiền gửi của khách hàng
	ChangeDerivativesLiabilities    *float64 `json:"change_derivatives_liabilities"`     // Biến động công cụ phái sinh và nợ tài chính khác
	ChangeTrustInvestments          *float64 `json:"change_trust_investments"`           // Biến động quỹ tín thác và đầu tư từ chính phủ và các bên khác
	ChangeIssuedDebt                *float64 `json:"change_issued_debt"`                 // Biến động công cụ nợ đã phát hành
	ChangeInterestFeePayables       *float64 `json:"change_interest_fee_payables"`       // Biến động lãi và phí phải trả
	ChangeOtherLiabilities          *float64 `json:"change_other_liabilities"`           // Biến động các khoản nợ hoạt động khác

	// Lưu chuyển tiền tệ từ hoạt động đầu tư
	NetCashInvesting                *float64 `json:"net_cash_investing"`                 // Lưu chuyển tiền thuần từ đầu tư
	PaymentsForFixedAssets          *float64 `json:"payments_for_fixed_assets"`          // Chi trả mua tài sản cố định
	ProceedsFromFixedAssets         *float64 `json:"proceeds_from_fixed_assets"`         // Thu nhập từ bán tài sản cố định
	PaymentsForAssetDisposals       *float64 `json:"payments_for_asset_disposals"`       // Chi trả thanh lý tài sản cố định
	InvestmentsInEntities           *float64 `json:"investments_in_entities"`            // Đầu tư vào tổ chức khác
	ProceedsFromInvestmentDisposals *float64 `json:"proceeds_from_investment_disposals"` // Thu nhập từ thanh lý đầu tư
	DividendsInterestReceived       *float64 `json:"dividends_interest_received"`        // Cổ tức và thu nhập lãi đã nhận

	// Lưu chuyển tiền tệ từ hoạt động tài chính
	NetCashFinancing          *float64 `json:"net_cash_financing"`           // Lưu chuyển tiền thuần từ tài chính
	ProceedsFromSharesCapital *float64 `json:"proceeds_from_shares_capital"` // Thu từ phát hành cổ phiếu và góp vốn
	ProceedsFromLongTermDebt  *float64 `json:"proceeds_from_long_term_debt"` // Thu từ phát hành nợ dài hạn
	PaymentsForLongTermDebt   *float64 `json:"payments_for_long_term_debt"`  // Chi trả nợ dài hạn
	DividendsPaid             *float64 `json:"dividends_paid"`               // Cổ tức đã chi trả

	// Tổng hợp lưu chuyển tiền tệ
	NetChangeCashEquivalents *float64 `json:"net_change_cash_equivalents"` // Biến động thuần tiền và tương đương tiền
	CashEquivalentsAtStart   *float64 `json:"cash_equivalents_at_start"`   // Tiền và tương đương tiền đầu kỳ
	ExchangeRateEffect       *float64 `json:"exchange_rate_effect"`        // Ảnh hưởng của biến động tỷ giá hối đoái
	CashEquivalentsAtEnd     *float64 `json:"cash_equivalents_at_end"`     // Tiền và tương đương tiền cuối kỳ

	// ================================================================= Balance Sheet =================================================================
	// Tổng tài sản
	TotalCompanyAssets *float64 `json:"total_company_assets"` // Tổng tài sản

	// Tiền mặt và tương đương tiền
	CashGoldSilver            *float64 `json:"cash_gold_silver"`            // Tiền mặt, vàng, bạc
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

	// fix =================================================================
	CustomerLoans          *float64 `json:"customer_loans"`           // Tên giữ nguyên
	ProvisionCustomerLoans *float64 `json:"provision_customer_loans"` // Trường mới thêm

	CorporateIncomeTaxPaidAgain              *float64 `json:"corporate_income_tax_paid_again"`              // Trường mới thêm
	PaymentsFromCreditInstitutionFunds       *float64 `json:"payments_from_credit_institution_funds"`       // Trường mới thêm
	ReceiptsFromBadDebtRecoveries            *float64 `json:"receipts_from_bad_debt_recoveries"`            // Trường mới thêm
	PaymentsForInvestmentProperties          *float64 `json:"payments_for_investment_properties"`           // Trường mới thêm
	ProceedsFromInvestmentPropertiesDisposal *float64 `json:"proceeds_from_investment_properties_disposal"` // Trường mới thêm
	PaymentsForInvestmentPropertiesDisposal  *float64 `json:"payments_for_investment_properties_disposal"`  // Trường mới thêm
	PaymentsForPurchaseOfTreasuryShares      *float64 `json:"payments_for_purchase_of_treasury_shares"`     // Trường mới thêm
	ProceedsFromSaleOfTreasuryShares         *float64 `json:"proceeds_from_sale_of_treasury_shares"`        // Trường mới thêm

	// IncomeStatement   IncomeStatement   `json:"income_statement" gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// CashFlowStatement CashflowStatement `json:"cash_flow" gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// BalanceSheet      BalanceSheet      `json:"balance_sheet" gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// FinancialAnalyst  FinancialAnalyzer `json:"financial_analyst" gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
