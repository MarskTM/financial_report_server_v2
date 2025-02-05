package service

import (
	"phenikaa/infrastructure"
	"phenikaa/model"

	"gorm.io/gorm"
)

type documentService struct {
	db *gorm.DB
}
type DocumentService interface {
	UploadFile(data model.Document) (*model.Document, error)
	DeleteFileReport(documentId int32) error
}

func (s *documentService) UploadFile(data model.Document) (*model.Document, error) {
	if err := s.db.Model(&model.Document{}).Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func NewDocumentService() *documentService {
	db := infrastructure.GetDB()
	return &documentService{
		db: db,
	}
}
