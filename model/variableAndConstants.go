package model

var (
	AccessTokenTime  int64 = 24
	RefreshTokenTime int64 = 72
	DefaultPassword        = "phenikaa@123"
)

var MapModelType = map[string]interface{}{
	"users":               []User{},
	"roles":               []Role{},
	"userRole":            []UserRole{},
	"profiles":            []Profile{},
	"companies":           []Company{},
	"userForgotPasswords": []UserForgotPassword{},
	"financialReports":    []FinancialReport{},
	"companyStakeholders": []CompanyStakeholder{},
	"userReports":         []UserReport{},
	"companyReports":      []CompanyReport{},
}

var MapAssociation = map[string]map[string]interface{}{ // Alown preload association 2 level model
	"users": {
		"UserRoles":      "",
		"UserRoles.Role": "",
	},
	"roles":     {},
	"userRoles": {},
	"profiles": {
		"User":                "",
		"User.UserRoles.Role": "",
		"Recruitment":         "",
		"InternShip":          "",
	},
	"financialReports": {},
	"companies": {
		"CompanyStakeholders": "",
	},
	"companyStakeholders": {},
	"userReports": {
		"Reports": "",
	},
    "companyReports": {
		"Reports": "",
	},
}
