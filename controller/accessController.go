package controller

import (
	"encoding/json"
	"fmt"
	"phenikaa/model"
	"phenikaa/service"
	"strings"

	// "strings"

	"github.com/go-chi/render"
	"gorm.io/gorm"

	"net/http"
)

type accessController struct {
	accessService service.AccessService
	userService   service.UserService
	db            *gorm.DB
}

type AccessController interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
}

// @Summary Login
// @Description Login
// @Tags Access
// @Accept json
// @Produce json
// @Param payload body model.LoginPayload true "Login"
// @Success 200 {object} Response
// @Router /login [post]
func (c *accessController) Login(w http.ResponseWriter, r *http.Request) {
	var res *Response
	var payload model.LoginPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	if check, err := c.userService.CheckCredentials(payload.Username, payload.Password); err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	} else if !check {
		InternalServerErrorResponse(w, r, fmt.Errorf("Credentials was not match, auth: %v", check))
		return
	}

	userInfo, err := c.userService.GetByUsername(payload.Username)
	if err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	tokenDetail, err := c.accessService.CreateToken(uint(userInfo.ID), userInfo.Role)
	if err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	// if err := c.accessService.CreateAuth(int(userInfo.ID), tokenDetail); err != nil {
	// 	InternalServerErrorResponse(w, r, err)
	// 	return
	// }

	userInfo.AccessToken = tokenDetail.AccessToken
	userInfo.RefreshToken = tokenDetail.RefreshToken
	fullDomain := r.Header.Get("Origin")
	errCookie := SaveHttpCookie(fullDomain, tokenDetail, w)
	if errCookie != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}
	res = &Response{
		Data:    userInfo,
		Success: true,
		Message: "Login success",
	}

	render.JSON(w, r, res)
	return
}

// @Summary Logout
// @Description Logout
// @Tags Access
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /logout [post]
func (c *accessController) Logout(w http.ResponseWriter, r *http.Request) {
	return
}

// @Summary Refresh
// @Description Refresh
// @Tags Access
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /refresh [post]
func (c *accessController) Refresh(w http.ResponseWriter, r *http.Request) {
	var res Response
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		BadRequestResponse(w, r, fmt.Errorf("Authorization header not found"))
		return
	}
	authorizationBearer := strings.Split(authorization, " ")[1]
	accessToken := strings.Split(authorizationBearer, ";")[0]
	accessClaims, errDecodeToken := GetAndDecodeToken(accessToken)
	if errDecodeToken != nil {
		UnauthorizedResponse(w, r, errDecodeToken)
		return
	}

	refreshToken := strings.Split(authorizationBearer, ";")[1]
	refreshClaims, errDecodeToken := GetAndDecodeToken(refreshToken)
	if errDecodeToken != nil {
		UnauthorizedResponse(w, r, errDecodeToken)
		return
	}

	accessUuid := accessClaims["access_uuid"].(string)
	refreshUuid := refreshClaims["refresh_uuid"].(string)
	userId := uint(accessClaims["user_id"].(float64))
	role := refreshClaims["role"].(string)
	username := refreshClaims["username"].(string)

	// Delete the previous Refresh Token
	deleteAccess, errDelete := c.accessService.DeleteAuth(accessUuid)
	if errDelete != nil || deleteAccess == 0 { // if any goes wrong
		ForbiddenResponse(w, r, errDelete)
	}

	deletedRefesh, errDelete := c.accessService.DeleteAuth(refreshUuid)
	if errDelete != nil || deletedRefesh == 0 { // if any goes wrong
		ForbiddenResponse(w, r, errDelete)
	}

	// Create new pairs of refresh and access tokens
	tokenDetail, err := c.accessService.CreateToken(userId, role)
	if err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	// Create new authorization
	if err := c.accessService.CreateAuth(int(userId), tokenDetail); err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	userInfo, err := c.userService.GetByUsername(username)
	if err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	// fullDomain := r.Header.Get("Origin")
	// errCookie := SaveHttpCookie(fullDomain, tokenDetail, w)
	// if errCookie != nil {
	// 	InternalServerErrorResponse(w, r, err)
	// 	return
	// }

	userInfo.AccessToken = tokenDetail.AccessToken
	userInfo.RefreshToken = tokenDetail.RefreshToken
	res = Response{
		Data:    userInfo,
		Success: true,
		Message: "Refresh success",
	}
	render.JSON(w, r, res)
	return
}

func NewAccessController() AccessController {
	return &accessController{
		accessService: service.NewAccessService(),
		userService:   service.NewUserService(),
	}
}
