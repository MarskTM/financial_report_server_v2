package controller

import (
	"encoding/json"
	"net/http"
	"phenikaa/model"
	"phenikaa/service"

	"github.com/go-chi/render"
)

type UserController interface {
	Register(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
	ChangePassowrd(w http.ResponseWriter, r *http.Request)
	ForgotPassword(w http.ResponseWriter, r *http.Request)
	CheckEmailExact(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userService service.UserService
}

// @Summary Register
// @Description Register
// @Tags Access
// @Accept json
// @Produce json
// @Param pauload body model.RegisterPayload true "UserRegister"
// @Success 200 {object} Response
// @Router /users/register [post]
func (c *userController) Register(w http.ResponseWriter, r *http.Request) {
	var res *Response
	var payload model.RegisterPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	newUser, err := c.userService.CreateUser(payload)
	if err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	res = &Response{
		Data:    newUser,
		Success: true,
		Message: "Register success",
	}
	render.JSON(w, r, res)
	return
}

// @Summary change password
// @Description change password
// @Tags Access
// @Accept json
// @Produce json
// @Authorization
// @Param pauload body string true "Username"
// @Success 200 {object} Response
func (c *userController) ChangePassowrd(w http.ResponseWriter, r *http.Request) {
	var res *Response
	var payload model.ChangePasswordPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	if err := c.userService.ChangePassword(payload); err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	res = &Response{
		Success: true,
		Message: "Change password success",
	}
	render.JSON(w, r, res)
	return
}

// @Summary Forgot password
// @Description Forgot password
// @Tags Access
// @Accept json
// @Produce json
// @Param pauload body model.ForgotPasswordPayload true "ForgotPassword"
// @Success 200 {object} Response
func (c *userController) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var res *Response
	var payload model.ForgotPasswordPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	if err := c.userService.ForgotPassword(payload); err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	res = &Response{
		Success: true,
		Message: "forgot password success",
	}
	render.JSON(w, r, res)
	return
}

// @Summary Reset password
// @Description Reset password
// @Tags Access
// @Accept json
// @Produce json
// @Param pauload body string true "Username"
// @Success 200 {object} Response
// @Router /users/reset-password [put]
func (c *userController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var res *Response
	var payload string

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	if err := c.userService.ResetPassword(payload); err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	res = &Response{
		Success: true,
		Message: "Reset password success",
	}
	render.JSON(w, r, res)
	return
}

// @Summary Check email exact
// @Description Check email exact
// @Tags Access
// @Accept json
// @Produce json
// @Param email query string true "Email"
// @Success 200 {object} Response
// @Router /users/check-email-exact [get]
func (c *userController) CheckEmailExact(w http.ResponseWriter, r *http.Request) {
	var res *Response
	var payload model.EmailForgotPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	if err := c.userService.CheckEmailExact(payload.Email); err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	res = &Response{
		Success: true,
		Message: "Check email exact success",
	}
	render.JSON(w, r, res)
	return
}

// @Summary Get all users
// @Description Get all users
// @Tags Admin
// @Produce json
func (c *userController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var res *Response
	var users []*model.UserSystemResponse

	users, err := c.userService.GetAllUsers()
	if err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	res = &Response{
		Data:    users,
		Success: true,
		Message: "Get all users success",
	}
	render.JSON(w, r, res)
}

func NewUserController() UserController {
	return &userController{
		userService: service.NewUserService(),
	}
}
