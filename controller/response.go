package controller

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set(http.StatusText(http.StatusBadRequest), err.Error())
	res := Response{
		Data:    nil,
		Success: false,
		Message: "Bad request. " + err.Error(),
	}
	render.JSON(w, r, res)
}

func InternalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set(http.StatusText(http.StatusInternalServerError), err.Error())
	res := Response{
		Data:    nil,
		Success: false,
		Message: "Internal Server Error. " + err.Error(),
	}
	render.JSON(w, r, res)
}

func UnauthorizedResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set(http.StatusText(http.StatusUnauthorized), err.Error())
	res := Response{
		Data:    nil,
		Success: false,
		Message: "Unauthorized. " + err.Error(),
	}
	render.JSON(w, r, res)
}

func ForbiddenResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusForbidden)
	w.Header().Set(http.StatusText(http.StatusForbidden), err.Error())
	res := Response{
		Data:    nil,
		Success: false,
		Message: "Forbidden. " + err.Error(),
	}
	render.JSON(w, r, res)
}
