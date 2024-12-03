package model

type BalanceSheet struct {
	ID int32 `json:"id" gorm:"primarykey"`
	// Tài sản ngắn hạn
	TotalAssets                       float64 `json:"total_assets" vi:"TỔNG TÀI SẢN"`
	CurrentAssets                     float64 `json:"current_assets" vi:"TÀI SẢN NGẮN HẠN"`
	ShortTermFinancialAssets          float64 `json:"short_term_financial_assets" vi:"Tài sản tài chính ngắn hạn"`
	CashAndCashEquivalents            float64 `json:"cash_and_cash_equivalents" vi:"Tiền và tương đương tiền"`
	Cash                              float64 `json:"cash" vi:"Tiền"`
	ShortTermInvestments              float64 `json:"short_term_investments" vi:"Các khoản tương đương tiền"`
	NetInvestments                    float64 `json:"net_investments" vi:"Giá trị thuần đầu tư tài sản tài chính ngắn hạn"`
	AssetsFVTPL                       float64 `json:"assets_fvtpl" vi:"Các tài sản tài chính ghi nhận thông qua lãi/lỗ (FVTPL)"`
	HTMInvestments                    float64 `json:"htm_investments" vi:"Các khoản đầu tư nắm giữ đến ngày đáo hạn (HTM)"`
	Loans                             float64 `json:"loans" vi:"Các khoản cho vay"`
	AFSInvestments                    float64 `json:"afs_investments" vi:"Các khoản tài chính sẵn sàng để bán (AFS)"`
	ProvisionForStockDepreciation     float64 `json:"provision_for_stock_depreciation" vi:"Dự phòng giảm giá chứng khoán kinh doanh"`
	TotalReceivables                  float64 `json:"total_receivables" vi:"Tổng các khoản phải thu"`
	ReceivablesFrom2016               float64 `json:"receivables_from_2016" vi:"Các khoản phải thu (từ 2016)"`
	ReceivablesCustomers              float64 `json:"receivables_customers" vi:"Phải thu khách hàng"`
	VATDeductible                     float64 `json:"vat_deductible" vi:"Thuế giá trị gia tăng được khấu trừ"`
	Prepayments                       float64 `json:"prepayments" vi:"Trả trước người bán"`
	InternalReceivables               float64 `json:"internal_receivables" vi:"Phải thu nội bộ"`
	ConstructionReceivables           float64 `json:"construction_receivables" vi:"Phải thu về XDCB"`
	OtherReceivables                  float64 `json:"other_receivables" vi:"Phải thu khác"`
	ProvisionForDoubtfulReceivables   float64 `json:"provision_for_doubtful_receivables" vi:"Dự phòng nợ khó đòi"`
	InventoryNet                      float64 `json:"inventory_net" vi:"Hàng tồn kho (Ròng)"`
	Inventory                         float64 `json:"inventory" vi:"Hàng tồn kho"`
	ProvisionForInventoryDepreciation float64 `json:"provision_for_inventory_depreciation" vi:"Dự phòng giảm giá HTK"`
	OtherCurrentAssets                float64 `json:"other_current_assets" vi:"Tài sản lưu động khác"`

	// Tài sản dài hạn
	LongTermAssets          float64 `json:"long_term_assets" vi:"TÀI SẢN DÀI HẠN"`
	LongTermFinancialAssets float64 `json:"long_term_financial_assets" vi:"Tài sản tài chính dài hạn"`
	LongTermReceivables     float64 `json:"long_term_receivables" vi:"Phải thu dài hạn"`
	Investments             float64 `json:"investments" vi:"Đầu tư dài hạn"`
	FixedAssetsNet          float64 `json:"fixed_assets_net" vi:"GTCL TSCĐ hữu hình"`
	FixedAssetsGross        float64 `json:"fixed_assets_gross" vi:"Nguyên giá TSCĐ hữu hình"`
	AccumulatedDepreciation float64 `json:"accumulated_depreciation" vi:"Khấu hao lũy kế TSCĐ hữu hình"`

	// Nợ phải trả
	TotalLiabilities        float64 `json:"total_liabilities" vi:"NỢ PHẢI TRẢ"`
	ShortTermLiabilities    float64 `json:"short_term_liabilities" vi:"Nợ ngắn hạn"`
	ShortTermFinancialDebts float64 `json:"short_term_financial_debts" vi:"Vay và nợ thuê tài chính ngắn hạn"`
	LongTermLiabilities     float64 `json:"long_term_liabilities" vi:"Nợ dài hạn"`

	// Vốn chủ sở hữu
	TotalEquity      float64 `json:"total_equity" vi:"VỐN CHỦ SỞ HỮU"`
	OwnerInvestment  float64 `json:"owner_investment" vi:"Vốn đầu tư của chủ sở hữu"`
	RetainedEarnings float64 `json:"retained_earnings" vi:"Lợi nhuận chưa phân phối"`

	// Tổng cộng nguồn vốn
	TotalCapital float64 `json:"total_capital" vi:"TỔNG CỘNG NGUỒN VỐN"`
}

