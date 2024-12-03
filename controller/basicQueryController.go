package controller

import (
	"encoding/json"
	"net/http"
	"phenikaa/model"
	"phenikaa/service"

	"github.com/go-chi/render"
)

type BasicQueryController interface {
	Upsert(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type basicQueryController struct {
	BasicQueryService service.BasicQueryService
}

// Upsert model to database
// @Summary Basic Query
// @Description Upsert model to database
// @Tags BasicQuery
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param basicQueryPayload body model.BasicQueryPayload true "BasicQueryPayload"
// @Success 200 {object} Response
// @Router /basic-query [post]
func (c *basicQueryController) Upsert(w http.ResponseWriter, r *http.Request) {
	var payload model.BasicQueryPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	temp, err := c.BasicQueryService.Upsert(payload)
	if err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	res := Response{
		Data:    temp,
		Success: true,
		Message: "Upsert " + payload.ModelType + "success",
	}
	render.JSON(w, r, res)
	return
}

// Delete model from database
// @Summary Basic Query
// @Description Delete model from database
// @Tags BasicQuery
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param basicQueryPayload body model.ListModelId true "BasicQueryPayload"
// @Success 200 {object} Response
// @Router /basic-query [delete]
func (c *basicQueryController) Delete(w http.ResponseWriter, r *http.Request) {
	var res Response
	var payload model.ListModelId
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	if err := c.BasicQueryService.Delete(payload); err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	res = Response{
		Data:    nil,
		Success: true,
		Message: "Delete " + payload.ModelType + " success",
	}
	render.JSON(w, r, res)
	return
}

func NewBasicQueryController() BasicQueryController {
	return &basicQueryController{
		BasicQueryService: service.NewBasicQueryService(),
	}
}
