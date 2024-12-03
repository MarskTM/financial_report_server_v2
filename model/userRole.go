package model

import "time"

type UserRole struct {
	ID     int32 `json:"id" gorm:"primaryKey"`
	UserID int32 `json:"user_id"`
	RoleID int32 `json:"role_id"`
	Active bool `json:"active"`

	Role *Role `json:"role" gorm:"foreignKey:RoleID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
	DeletedAt time.Time `json:"-" swaggerignore:"true"`
}
