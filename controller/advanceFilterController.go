package controller

import (
	"encoding/json"
	"net/http"
	"phenikaa/model"
	"phenikaa/service"

	"github.com/go-chi/render"
)

type AdvanceFilterController interface {
	Filter(w http.ResponseWriter, r *http.Request)
}

type advanceFilterController struct {
	AdvanceFilterService service.AdvanceFilter
}

// AdvanceFilterController
// @Summary Advance Filter
// @Description Advance Filter for all model
// @Tags Advance Filter
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param advanceFilterPayload body model.AdvanceFilterPayload true "Advance Filter Payload"
// @Success 200  {object} Response
// @Router /advance-filter [post]
func (c *advanceFilterController) Filter(w http.ResponseWriter, r *http.Request) {
	var res Response
	var payload model.AdvanceFilterPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	temp, err := c.AdvanceFilterService.Filter(payload)
	if err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	res = Response{
		Data:    temp,
		Success: true,
		Message: "Get " + payload.ModelType + " success",
	}
	render.JSON(w, r, res)
}

func NewAdvanceFilterController() AdvanceFilterController {
	return &advanceFilterController{
		AdvanceFilterService: service.NewAdvanceFilterController(),
	}
}
