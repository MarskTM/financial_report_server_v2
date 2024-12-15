package model

type UserReport struct {
	ID         int32  `json:"id" gorm:"primarykey"`
	ProfileID  int32  `json:"profile_id"`
	ReportID   int32  `json:"report_id"`
	DocumentID int32  `json:"document_id"`
	Quater     string `json:"quater"` // Quý báo cáo.

	Report   FinancialReport `json:"report" gorm:"foreignKey:ReportID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Document Document        `json:"document"  gorm:"foreignKey:DocumentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
