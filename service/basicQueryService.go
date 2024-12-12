package service

import (
	"fmt"
	"phenikaa/infrastructure"
	"phenikaa/model"
	"phenikaa/utils"
	"reflect"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

type BasicQueryService interface {
	Upsert(payload model.BasicQueryPayload) (interface{}, error)
	Delete(payload model.ListModelId) error
}

type basicQueryService struct {
	emailService EmailService
}

func (s *basicQueryService) Upsert(payload model.BasicQueryPayload) (interface{}, error) {
	var db = infrastructure.GetDB()
	var tableName = strcase.ToSnake(payload.ModelType)
	var modelType = model.MapModelType[payload.ModelType]

	// Get the list id exists in the database
	var listModelId = make([]uint, 0)
	if err := db.Model(modelType).Pluck("id", &listModelId).Error; err != nil {
		return nil, fmt.Errorf("get list id error: %v", err)
	}

	// Setval the max id in the database
	var maxModelId uint
	// queryGetMaxId := "SELECT setval('" + tableName + "_id_seq', (SELECT MAX(id) FROM " + tableName + ")+1);"
	// if err := db.Model(modelType).Raw(queryGetMaxId).Scan(&maxModelId).Error; err != nil {
	// 	return nil, fmt.Errorf("set max id error: %v", err)
	// }

	if err := db.Table(tableName).Select("COALESCE(MAX(id), 1)").Scan(&maxModelId).Error; err != nil {
		return nil, fmt.Errorf("set max id error: %v", err)
	}

	// Upsert multiple
	var listModelCreate []map[string]interface{}
	var listModelUpdate []map[string]interface{}
	
	if reflect.TypeOf(payload.Data).Kind() == reflect.Slice || reflect.TypeOf(payload.Data).Elem().Kind() == reflect.Slice {
		newInsertID := maxModelId
		for _, data := range payload.Data.([]interface{}) {
			newInsertID++

			data := data.(map[string]interface{})
			if data["id"] == nil || data["id"].(uint) == 0 {
				data["id"] = newInsertID
				listModelCreate = append(listModelCreate, data)
				continue
			}

			if ok, _ := utils.InArray(data["id"].(uint), listModelId); ok {
				listModelUpdate = append(listModelUpdate, data)
			}
		}

		if err := db.Transaction(func(tx *gorm.DB) error {
			if len(listModelCreate) > 0 {
				if err := tx.Model(modelType).CreateInBatches(&listModelCreate, 1000).Error; err != nil {
					return fmt.Errorf("create error: %v", err)
				}
			}

			if len(listModelUpdate) > 0 {
				if err := tx.Model(modelType).Updates(listModelUpdate).Error; err != nil {
					return fmt.Errorf("update error: %v", err)
				}
			}
			return nil
		}); err != nil {
			return nil, fmt.Errorf("upsert error: %v", err)
		}
		goto End
	}

	// Upsert single
	if payload.Data == nil {
		return nil, fmt.Errorf("data is nil cannot upsert")
	}

	if payload.Data.(map[string]interface{})["id"] == nil {
		payload.Data.(map[string]interface{})["id"] = maxModelId + 1
		if err := db.Model(modelType).Create(payload.Data.(map[string]interface{})).Error; err != nil {
			return nil, fmt.Errorf("create error: %v", err)
		}

		goto End
	} else {
		var modelId uint = uint(payload.Data.(map[string]interface{})["id"].(float64))
		if ok, _ := utils.InArray(modelId, listModelId); ok {
			if err := db.Model(modelType).Where("id = (?)", modelId).Updates(payload.Data.(map[string]interface{})).Error; err != nil {
				return nil, fmt.Errorf("update error: %v", err)
			}
		}
	}

	if len(payload.ModelType) == 0 {
		return nil, nil
	}

End:
	return payload.Data, nil
}

func (s *basicQueryService) Delete(payload model.ListModelId) error {
	var db = infrastructure.GetDB()
	var modelType = model.MapModelType[payload.ModelType]
	if err := db.Where("id IN (?)", payload.ID).Delete(modelType).Error; err != nil {
		return fmt.Errorf("Delete error: %v", err)
	}

	return nil
}

func NewBasicQueryService() BasicQueryService {
	return &basicQueryService{
		emailService: NewEmailService(),
	}
}
