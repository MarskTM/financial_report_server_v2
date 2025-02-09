package controller

import (
	"encoding/json"
	"net/http"
	"phenikaa/model"
	"phenikaa/service"

	"github.com/go-chi/render"
	"github.com/golang/glog"
)

type TidingController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type tidingController struct {
	tidingService service.TidingService
}

func (c *tidingController) GetAll(w http.ResponseWriter, r *http.Request) {
	tidings, err := c.tidingService.GetAll()
	if err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	// Send the response
	res := &Response{
		Data:    tidings,
		Success: true,
		Message: "Get all tiding success",
	}
	glog.V(3).Infof("+ Get all tiding response: %v", res)
	render.JSON(w, r, res)
}

func (c *tidingController) Create(w http.ResponseWriter, r *http.Request) {
	request := model.Tiding{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	newData, err := c.tidingService.Create(request)
	if err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	// Send the response
	res := &Response{
		Data:    newData,
		Success: true,
		Message: "Create tiding success",
	}
	glog.V(3).Infof("+ Create tiding response: %v", res)
	render.JSON(w, r, res)
}

func (c *tidingController) Update(w http.ResponseWriter, r *http.Request) {
	request := []model.Tiding{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	err := c.tidingService.Update(request)
	if err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}
	// Send the response
	res := &Response{
		Data:    nil,
		Success: true,
		Message: "Update tiding success",
	}
	glog.V(3).Infof("+ Update tiding response: %v", res)
	render.JSON(w, r, res)
}

func (c *tidingController) Delete(w http.ResponseWriter, r *http.Request) {
	request := model.Tiding{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		BadRequestResponse(w, r, err)
		return
	}
}

func NewTidingController() TidingController {
	tidingService := service.NewTidingService()
	return &tidingController{tidingService}
}
