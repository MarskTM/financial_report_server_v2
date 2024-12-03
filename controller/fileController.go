package controller

import (
	"fmt"
	"log"
	"net/http"
	"phenikaa/service"
)

type DocumentController interface {
	ImportReportData(w http.ResponseWriter, r *http.Request)
}

type documentController struct {
	BasicQueryService service.BasicQueryService
}

func (d *documentController) ImportReportData(w http.ResponseWriter, r *http.Request) {
	// Implement the logic for importing data from a file
	// var res *Response

	// Giới hạn kích thước tối đa của form, ví dụ 20 MB
	err := r.ParseMultipartForm(10 << 40)
	if err != nil {
		http.Error(w, "Form quá lớn", http.StatusBadRequest)
		return
	}

	// Lấy file "profile_picture" từ form
	file, fileHeader, err := r.FormFile("financial")
	if err != nil {
		BadRequestResponse(w, r, fmt.Errorf("Không thể lấy file từ form: %v", err))
		return
	}
	defer file.Close()
	log.Println("FileName:", fileHeader.Filename)

	// 1. Save file

	// 2. Open file and read data

	// 3. Import data to database

}

func NewDocumentController() DocumentController {
	return &documentController{
		BasicQueryService: service.NewBasicQueryService(),
	}
}
