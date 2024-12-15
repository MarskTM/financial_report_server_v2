package service

import (
	"phenikaa/infrastructure"
	"phenikaa/model"

	"gorm.io/gorm"
)

type DocumentService interface {
	ImportBalanceSheet(data model.BalanceSheet) (*model.BalanceSheet, error) 
	ImportIncomeStatement(data model.IncomeStatement) (*model.IncomeStatement, error)
	ImportCashFlowStatement(data model.CashflowStatement) (*model.CashflowStatement, error)
}

func (s *documentService) ImportBalanceSheet(data model.BalanceSheet) (*model.BalanceSheet, error) {
	err := s.db.Model(&model.BalanceSheet{}).Create(&data).Error
    if err!= nil {
        return nil, err
    }
    return &data, nil
}

func (s *documentService) ImportIncomeStatement(data model.IncomeStatement) (*model.IncomeStatement, error) {
	err := s.db.Model(&model.IncomeStatement{}).Create(&data).Error
    if err!= nil {
        return nil, err
    }
    return &data, nil
}

func (s *documentService) ImportCashFlowStatement(data model.CashflowStatement) (*model.CashflowStatement, error) {
	err := s.db.Model(&model.CashflowStatement{}).Create(&data).Error
    if err!= nil {
        return nil, err
    }
    return &data, nil
}

type documentService struct {
	db *gorm.DB
}

func NewDocumentService() *documentService {
	db := infrastructure.GetDB()
	return &documentService{
		db: db,
	}
}
