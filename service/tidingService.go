package service

import (
	"phenikaa/infrastructure"
	"phenikaa/model"

	"gorm.io/gorm"
)

type TidingService interface {
	GetAll() ([]model.Tiding, error)
	Create(new model.Tiding) (*model.Tiding, error)
	Update(new model.Tiding) (*model.Tiding, error)
	Delete(id int32) error
}

type tidingService struct {
	db *gorm.DB
}

func (s *tidingService) GetAll() ([]model.Tiding, error) {
	var tidings []model.Tiding
	if err := s.db.Where("parent_id IS NULL").Preload("SubTidings").Find(&tidings).Error; err != nil {
		return nil, err
	}

	return tidings, nil
}

func (s *tidingService) Create(new model.Tiding) (*model.Tiding, error) {
	err := s.db.Debug().Transaction(func(tx *gorm.DB) error {
		// Ensure ParentID is nil if not provided
		if new.ParentID != nil && *new.ParentID == 0 {
			new.ParentID = nil
		}

		// Create the parent Tiding
		new.State = true
		if err := tx.Create(&new).Error; err != nil {
			return err
		}

		// Save and associate SubTidings if any
		if len(new.SubTidings) > 0 {
			for i := range new.SubTidings {
				new.SubTidings[i].ParentID = &new.ID
			}
			if err := tx.Model(&new).Association("SubTidings").Replace(new.SubTidings); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &new, nil
}

func (s *tidingService) Update(data model.Tiding) (*model.Tiding, error) {
	err := s.db.Debug().Transaction(func(tx *gorm.DB) error {
		parent := data
		listChildren := data.SubTidings

		if err := tx.Where("id = ?", parent.ID).Model(&model.Tiding{}).Updates(&parent).Error; err != nil {
			return err
		}

		// Upsert SubTidings
		for _, child := range listChildren {
			var existingChild model.Tiding
			if err := tx.Where("title = ? AND parent_id = ?", child.Title, parent.ID).First(&existingChild).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// Create new SubTiding if not found
					child.ParentID = &parent.ID
					if err := tx.Create(&child).Error; err != nil {
						return err
					}
				} else {
					return err
				}
			} else {
				// Update existing SubTiding
				child.ID = existingChild.ID
				child.ParentID = &parent.ID
				if err := tx.Model(&existingChild).Updates(&child).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *tidingService) Delete(id int32) error {
	return s.db.Model(&model.Tiding{}).Where("id = ?", id).Delete(&model.Tiding{}).Error
}

func NewTidingService() TidingService {
	db := infrastructure.GetDB()
	return &tidingService{db: db}
}
