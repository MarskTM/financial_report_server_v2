package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       int32  `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"type:varchar(100);unique"`
	Password string `json:"password"`

	UserRoles *UserRole `json:"user_roles" gorm:"foreignKey:UserID"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}

type UserForgotPassword struct {
	ID        int32          `json:"id" gorm:"primary_key"`
	UserId    int32          `json:"user_id" gorm:"unique"`
	FogotCode string         `json:"fogot_code"`
	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}

type UserResponse struct {
	ID           int32    `json:"id"`
	Role         string   `json:"role"`
	Username     string   `json:"username"`
	FullName     string   `json:"fullname"`
	Profile      *Profile `json:"profile"`
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
}
