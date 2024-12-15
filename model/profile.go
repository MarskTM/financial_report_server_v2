package model

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID         int32  `json:"id" gorm:"primaryKey"`
	UserID     int32  `json:"user_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Alias      string `json:"alias"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Birthdate  string `json:"birthdate"`
	OtherLinks string `json:"other_links"`
	Address    string `json:"address"`

	User *User `json:"user" gorm:"foreignKey:UserID"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}
