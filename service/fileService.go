package service

import (
	"phenikaa/infrastructure"
	"phenikaa/model"

	"github.com/golang/glog"
	"gorm.io/gorm"
)

type documentService struct {
	db *gorm.DB
}
type DocumentService interface {
	UploadFile(data model.Document) (*model.Document, error)
	GetFinancialReportByProfileId(profileId int32) (*model.FinancialReport, error)
	DeleteFileReport(documentId int32) error
}

func (s *documentService) UploadFile(data model.Document) (*model.Document, error) {
	if err := s.db.Model(&model.Document{}).Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *documentService) DeleteFileReport(reportID int32) error {
	glog.V(1).Infof("DeleteFileReport: %v", reportID)
	err := s.db.Debug().Transaction(func(tx *gorm.DB) error {
		var userReport model.UserReport
		if err := tx.Model(&model.UserReport{}).Where("id = ?", reportID).Find(&userReport).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.Document{}).Where("id = ?", userReport.DocumentID).Delete(&model.Document{}).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.FinancialReport{}).Where("user_report_id = ?", userReport.ProfileID).Delete(&model.FinancialReport{}).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.UserReport{}).Where("id = ?", userReport.ID).Delete(&model.UserReport{}).Error; err != nil {
			return err
		}

		return nil
	})

	// commit transaction
	if err != nil {
		return err
	}

	return nil
}

func (s *documentService) GetFinancialReportByProfileId(profileId int32) (*model.FinancialReport, error) {
	return nil, nil
}

func NewDocumentService() *documentService {
	db := infrastructure.GetDB()
	return &documentService{
		db: db,
	}
}
