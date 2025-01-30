package model

import "github.com/lib/pq"

type AdvanceFilterPayload struct {
	ModelType         string         `json:"modelType"`
	IgnoreAssociation []string       `json:"ignoreAssociation"`
	Page              int            `json:"page"`
	PageSize          int            `json:"pageSize"`
	IsPaginateDB      bool           `json:"isPaginateDB"`
	QuerySearch       string         `json:"querySearch"`
	SelectColumn      pq.StringArray `json:"selectColumn"`
	Sort              string         `json:"sort"`
}

type BasicQueryPayload struct {
	ModelType string      `json:"modelType"`
	Data      interface{} `json:"data"`
}

type ListModelId struct {
	ID        []uint `gorm:"column:id"`
	ModelType string `json:"modelType"`
}

// TokenDetail details for token authentication
type TokenDetail struct {
	Username     string
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

// AccessDetail access detail only from token
type AccessDetail struct {
	AccessUUID string
	UserID     int
}

// Payload for authentication
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterPayload struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Birthday  string `json:"birthday"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ChangePasswordPayload struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type ForgotPasswordPayload struct {
	FogortCode  string `json:"forgotCode"`
	NewPassword string `json:"newPassword"`
}

type EmailForgotPayload struct {
	Email string `json:"email"`
}

type UpdateUserStatePayload struct {
	ID       int32  `json:"id"`
	IsActive bool   `json:"isActive"`
	Role     string `json:"role"`
}
