package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/lithammer/shortuuid"
)

// GenerateKey using in set keys
func GenCode() string {
	id := shortuuid.New()
	return strings.ToUpper(id[0:10])
}

// PatternGet using in get keys
func PatternGet(id uint) string {
	return strconv.Itoa(int(id)) + "-:--*"
}

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// GenerateKey random password
func GeneratePasswordKey(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buffer)[:length], nil
}

func MapDataToStruct(data map[string]interface{}, target interface{}, tag string) error {
	// Lấy giá trị và kiểu của struct
	val := reflect.ValueOf(target).Elem()
	typ := val.Type()

	// Duyệt qua từng trường trong struct
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		viTag := field.Tag.Get(tag) // Lấy giá trị của tag vi

		// Kiểm tra nếu tag vi có trong data
		if value, exists := data[viTag]; exists {
			fieldValue := val.FieldByName(field.Name)

			// Gán giá trị nếu trường hợp hợp lệ và có thể gán
			if fieldValue.IsValid() && fieldValue.CanSet() {
				valToSet := reflect.ValueOf(value)
				if valToSet.Type().ConvertibleTo(fieldValue.Type()) {
					fieldValue.Set(valToSet.Convert(fieldValue.Type()))
				}
			}
		}
	}

	return nil
}

// ConvertToInt32 nhận một giá trị bất kỳ và cố gắng chuyển đổi nó về kiểu int32
func ConvertToInt32(value interface{}) (int32, error) {
	switch v := value.(type) {
	case int:
		return int32(v), nil
	case int32:
		return v, nil
	case int64:
		if v > int64(int32(^uint32(0)>>1)) || v < int64(-int32(^uint32(0)>>1)-1) {
			return 0, errors.New("value out of range for int32")
		}
		return int32(v), nil
	case int8:
		return int32(v), nil
	case float32:
		if v > float32(int32(^uint32(0)>>1)) || v < float32(-int32(^uint32(0)>>1)-1) {
			return 0, errors.New("value out of range for int32")
		}
		return int32(v), nil
	case float64:
		if v > float64(int32(^uint32(0)>>1)) || v < float64(-int32(^uint32(0)>>1)-1) {
			return 0, errors.New("value out of range for int32")
		}
		return int32(v), nil
	default:
		return 0, fmt.Errorf("unsupported type: %s", reflect.TypeOf(value).String())
	}
}
