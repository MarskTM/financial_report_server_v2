package service

import (
	"fmt"
	"phenikaa/infrastructure"
	"phenikaa/model"
)

type AdvanceFilter interface {
	Filter(payload model.AdvanceFilterPayload) (interface{}, error)
}

type advanceFilterService struct{}

func (s *advanceFilterService) Filter(payload model.AdvanceFilterPayload) (interface{}, error) {
	var db = infrastructure.GetDB()
	var modelType = model.MapModelType[payload.ModelType]
	var modelPreload = model.MapAssociation[payload.ModelType]
	if modelType == nil {
		return nil, fmt.Errorf("model type not found")
	}

	var query = "id > 0 AND " + payload.QuerySerch
	if payload.IsPaginateDB {
		db = db.Limit(payload.PageSize).Offset((payload.Page - 1) * payload.PageSize) // This offset to calculate the offset of the first row returned
	}

	if len(payload.SelectColumn) > 0 {
		db = db.Select(payload.SelectColumn)
	}
	if !payload.IgnoreAssociation {
		for model, condition := range modelPreload {
			db = db.Preload(model, condition)
		}
	}

	if err := db.Model(modelType).Where(query).Find(&modelType).Error; err != nil {
		return nil, err
	}
	return &modelType, nil
}

func NewAdvanceFilterController() AdvanceFilter {
	return &advanceFilterService{}
}
