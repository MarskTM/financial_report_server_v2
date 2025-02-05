package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"phenikaa/model"
	"phenikaa/service"
	"strings"

	"github.com/go-chi/render"
	"github.com/golang/glog"
)

type DocumentController interface {
	ImportReportData(w http.ResponseWriter, r *http.Request)
	ExportReportData(w http.ResponseWriter, r *http.Request)
}

type documentController struct {
	BasicQueryService service.BasicQueryService
	DocumentService   service.DocumentService
}

func (c *documentController) ImportReportData(w http.ResponseWriter, r *http.Request) {
	// Implement the logic for importing data from a file
	// var res *Response

	// Giới hạn kích thước tối đa của form, ví dụ 20 MB
	err := r.ParseMultipartForm(10 << 40)
	if err != nil {
		http.Error(w, "Form quá lớn", http.StatusBadRequest)
		return
	}

	// Lấy thông tin từ form
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		BadRequestResponse(w, r, fmt.Errorf("Không thể lấy file từ form: %v", err))
		return
	}
	defer file.Close()
	log.Println("FileName:", fileHeader.Filename)

	// 1. Save file
	fileName := fileHeader.Filename
	cdnDir := "./cdn"

	err = os.MkdirAll(cdnDir, os.ModePerm)
	if err != nil {
		// Xử lý lỗi
		return
	}
	filePath := filepath.Join(cdnDir, fileName)

	// Kiểm tra file đã tồn tại
	if _, err := os.Stat(filePath); err == nil { // File đã tồn tại
		ext := filepath.Ext(fileName)                 // Lấy phần mở rộng của file
		baseName := strings.TrimSuffix(fileName, ext) // Lấy tên file không có phần mở rộng

		i := 1
		for {
			newFileName := fmt.Sprintf("%s_%d%s", baseName, i, ext) // Tạo tên mới với tiền tố số
			newFilePath := filepath.Join(cdnDir, newFileName)
			if _, err := os.Stat(newFilePath); os.IsNotExist(err) { // Kiểm tra tên mới đã tồn tại chưa
				filePath = newFilePath // Sử dụng tên mới
				fileName = newFileName
				break
			}
			i++ // Tăng tiền tố số và thử lại
		}
	}

	// 2. Open file and read/write data
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		// Xử lý lỗi
		return
	}

	err = os.WriteFile(filePath, fileBytes, os.ModePerm)
	if err != nil {
		// Xử lý lỗi
		return
	}

	// 3. Import data to database
	log.Println("File saved to:", filePath) // In đường dẫn đầy đủ
	newData := model.Document{
		Title: fileName,
		Cdn:   filePath,
	}
	newDocument, err := c.DocumentService.UploadFile(newData)
	if err != nil {
		InternalServerErrorResponse(w, r, fmt.Errorf("không thể tải file lên cdn: %v", err))
		return
	}
	glog.V(3).Info("=======> data saved:", newDocument)

	res := &Response{
		Data:    newDocument,
		Success: true,
		Message: "Login success",
	}
	render.JSON(w, r, res)
	// ... (các xử lý khác)
}

func (c *documentController) ExportReportData(w http.ResponseWriter, r *http.Request) {

}

func NewDocumentController() DocumentController {
	return &documentController{
		BasicQueryService: service.NewBasicQueryService(),
		DocumentService:   service.NewDocumentService(),
	}
}
