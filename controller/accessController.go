package controller

import (
	"encoding/json"
	"fmt"
	"phenikaa/model"
	"phenikaa/service"
	"strings"

	// "strings"

	"github.com/go-chi/render"
	"github.com/golang/glog"
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
	glog.V(3).Infof("+ Login() request: %v", payload)

	// 1. Check Credentials
	userInfo, err := c.userService.CheckCredentials(payload.Username, payload.Password)
	if err != nil {
		errSystem := fmt.Errorf("- Login() - err: %v", err)
		InternalServerErrorResponse(w, r, errSystem)
		return
	}

	// 2. Generate the authen token
	tokenDetail, err := c.accessService.CreateToken(uint(userInfo.ID), userInfo.Role)
	if err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}
	userInfo.AccessToken = tokenDetail.AccessToken
	userInfo.RefreshToken = tokenDetail.RefreshToken

	// 3. Cache Access Token
	// if err := c.accessService.CreateAuth(int(userInfo.ID), tokenDetail); err != nil {
	// 	InternalServerErrorResponse(w, r, err)
	// 	return
	// }

	// 4. Save Access Token in the HTTP Cookie
	fullDomain := r.Header.Get("Origin")
	errCookie := SaveHttpCookie(fullDomain, tokenDetail, w)
	if errCookie != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	// 5. Send the response
	res = &Response{
		Data:    userInfo,
		Success: true,
		Message: "Login success",
	}
	glog.V(3).Infof("+ Login() response: %v", payload)
	render.JSON(w, r, res)
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
		err := fmt.Errorf("- Refresh(): authorization header not found")

		glog.V(3).Info(err)
		BadRequestResponse(w, r, err)
		return
	}
	glog.V(3).Infof("+ Refresh() request: %v", authorization)

	// 1. Get authorization information
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

	// accessUuid := accessClaims["access_uuid"].(string)
	// refreshUuid := refreshClaims["refresh_uuid"].(string)
	userId := uint(accessClaims["user_id"].(float64))
	role := refreshClaims["role"].(string)
	username := refreshClaims["username"].(string)

	//2. Create new pairs of refresh and access tokens
	tokenDetail, err := c.accessService.CreateToken(userId, role)
	if err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	// 3. Refresh cached token
	// Delete the previous Refresh Token
	// deleteAccess, errDelete := c.accessService.DeleteAuth(accessUuid)
	// if errDelete != nil || deleteAccess == 0 { // if any goes wrong
	// 	ForbiddenResponse(w, r, errDelete)
	// }

	// deletedRefesh, errDelete := c.accessService.DeleteAuth(refreshUuid)
	// if errDelete != nil || deletedRefesh == 0 { // if any goes wrong
	// 	ForbiddenResponse(w, r, errDelete)
	// }

	// if err := c.accessService.CreateAuth(int(userId), tokenDetail); err != nil {
	// 	InternalServerErrorResponse(w, r, err)
	// 	return
	// }

	// 4. Save Access Token in the HTTP Cookie
	userInfo, err := c.userService.GetByUsername(username)
	if err != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}
	userInfo.AccessToken = tokenDetail.AccessToken
	userInfo.RefreshToken = tokenDetail.RefreshToken

	// 5. Save Access Token in the HTTP Cookie
	fullDomain := r.Header.Get("Origin")
	errCookie := SaveHttpCookie(fullDomain, tokenDetail, w)
	if errCookie != nil {
		InternalServerErrorResponse(w, r, err)
		return
	}

	// 6. Send the response
	res = Response{
		Data:    userInfo,
		Success: true,
		Message: "Refresh success",
	}
	render.JSON(w, r, res)
    glog.V(3).Infof("+ Refresh() response: %v", authorization)
}

func NewAccessController() AccessController {
	return &accessController{
		accessService: service.NewAccessService(),
		userService:   service.NewUserService(),
	}
}
