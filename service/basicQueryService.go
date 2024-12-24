package service

import (
	"fmt"
	"phenikaa/infrastructure"
	"phenikaa/model"
	"phenikaa/utils"
	"reflect"

	"github.com/golang/glog"
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
	var listModelId = make([]int32, 0)
	if err := db.Model(modelType).Pluck("id", &listModelId).Error; err != nil {
		return nil, fmt.Errorf("get list id error: %v", err)
	}

	// Setval the max id in the database
	var maxModelId int32
	// queryGetMaxId := "SELECT setval('" + tableName + "_id_seq', (SELECT MAX(id) FROM " + tableName + ")+1);"
	// if err := db.Model(modelType).Raw(queryGetMaxId).Scan(&maxModelId).Error; err != nil {
	// 	return nil, fmt.Errorf("set max id error: %v", err)
	// }

	if err := db.Table(tableName).Select("COALESCE(MAX(id), 0)").Scan(&maxModelId).Error; err != nil {
		return nil, fmt.Errorf("set max id error: %v", err)
	}

	// Upsert multiple
	listModelCreate := []map[string]interface{}{}
	listModelUpdate := []map[string]interface{}{}

	if reflect.TypeOf(payload.Data).Kind() == reflect.Slice || reflect.TypeOf(payload.Data).Elem().Kind() == reflect.Slice {
		// log.Println("upsert ============================= ", payload.Data)
		newInsertID := maxModelId

		for _, dataTemp := range payload.Data.([]interface{}) {
			newInsertID += 1

			data := dataTemp.(map[string]interface{})

			// // log.Println("upsert1 ============================= ", data)
			// id, hasID := data["id"]
			// if !hasID {
			// 	if id, ok := id.(int32); ok && id == 0 {
			// 		// log.Println("upsert1.2 ============================= ", id == int32(0))

			// 		data["id"] = newInsertID
			// 		listModelCreate = append(listModelCreate, data)
			// 		continue
			// 	} else {
			// 		// log.Println("upsert1.3 ============================= ")
			// 		if ok, _ := utils.InArray(id, listModelId); ok {
			// 			listModelUpdate = append(listModelUpdate, data)
			// 		}
			// 	}
			// } else {
			// 	// log.Println("upsert1.4 ============================= ", newInsertID)
			// 	data["id"] = newInsertID
			// 	listModelCreate = append(listModelCreate, data)
			// }

			// Tối ưu xử lý id và thêm vào danh sách
			if idValue, hasID := data["id"]; hasID {
				// Trường hợp đã có ID
				id, err := utils.ConvertToInt32(idValue) // Sử dụng hàm ConvertToInt32
				if err == nil && id != 0 {
					// Nếu ID hợp lệ (int32) và khác 0
					if exists, _ := utils.InArray(id, listModelId); exists {
						// Nếu ID đã tồn tại trong listModelId, thêm vào danh sách cập nhật
						listModelUpdate = append(listModelUpdate, data)
					} else {
						// Nếu ID không tồn tại, thêm vào danh sách tạo mới
						data["id"] = newInsertID
						listModelCreate = append(listModelCreate, data)
					}
				} else {
					// Trường hợp ID không hợp lệ hoặc bằng 0, tạo mới
					data["id"] = newInsertID
					listModelCreate = append(listModelCreate, data)
				}
			} else {
				// Trường hợp không có ID, tạo mới
				data["id"] = newInsertID
				listModelCreate = append(listModelCreate, data)
			}
		}
		glog.V(3).Infof("+ listModelCreate: %v", listModelCreate)
		glog.V(3).Infof("+ listModelUpdate: %v", listModelUpdate)

		if err := db.Transaction(func(tx *gorm.DB) error {
			if len(listModelCreate) > 0 {
				// if err := tx.Debug().Model(modelType).CreateInBatches(listModelCreate, 1000).Error; err != nil {
				if err := tx.Debug().Table(tableName).CreateInBatches(listModelCreate, 1000).Error; err != nil {
					return fmt.Errorf("create error: %v", err)
				}
			}

			if len(listModelUpdate) > 0 {
				if err := tx.Debug().Table(tableName).Updates(listModelUpdate).Error; err != nil {
					return fmt.Errorf("update error: %v", err)
				}
			}
			return nil
		}); err != nil {
			errTransaction := fmt.Errorf("upsert error: %v", err)
			glog.Error(errTransaction)

			return nil, errTransaction
		}

		return payload.Data, nil
	}

	// Upsert single
	if payload.Data == nil {
		return nil, fmt.Errorf("data is nil cannot upsert")
	}

	// =========================================================== Tối ưu hàm kiểm tra và xử lý `id` =================================================================
	if dataMap, ok := payload.Data.(map[string]interface{}); ok {
		if idValue, hasID := dataMap["id"]; hasID {
			// Chuyển đổi ID về int32 do mặc dịnh khi decode interface sẽ tùy vào phiên bản mà có kiểu dữ liệu id: float64 hoặc int32
			id, err := utils.ConvertToInt32(idValue)
			if err != nil {
				return nil, fmt.Errorf("invalid id type, expected int32 but got: %T, error: %v", idValue, err)
			}

			glog.V(3).Infof("upsert2.11 ============================= id: %v", id)

			if id == 0 {
				// ID bằng 0, tạo mới với ID tăng
				newID := maxModelId + 1
				dataMap["id"] = newID

				if err := db.Debug().Table(tableName).Create(dataMap).Error; err != nil {
					return nil, fmt.Errorf("create error: %v", err)
				}
				goto End
			}

			// Kiểm tra nếu ID đã tồn tại
			if exists, _ := utils.InArray(id, listModelId); exists {
				if err := db.Debug().Table(tableName).Where("id = ?", id).Updates(dataMap).Error; err != nil {
					return nil, fmt.Errorf("update error: %v", err)
				}
				goto End
			}

			// Nếu ID không tồn tại trong danh sách
			return nil, fmt.Errorf("ID %v does not exist in the listModelId", id)
		} else {
			// Nếu không có ID, tạo mới với ID tăng
			newID := maxModelId + 1
			dataMap["id"] = newID

			if err := db.Debug().Table(tableName).Create(dataMap).Error; err != nil {
				return nil, fmt.Errorf("create error: %v", err)
			}
			goto End
		}
	} else {
		return nil, fmt.Errorf("payload.Data is not a valid map[string]interface{}")
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
