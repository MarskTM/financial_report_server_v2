package controller

import (
	"net/http"
	"phenikaa/infrastructure"
	"phenikaa/model"

	"gorm.io/gorm"
)

type SeedController interface {
	SeedDatabase(w http.ResponseWriter, r *http.Request)
}

type seedController struct {
	db *gorm.DB
}

func (c *seedController) SeedDatabase(w http.ResponseWriter, r *http.Request) {
	user := model.User {
		ID: 0,
	}
	
	if err := c.db.Debug().Clauses().Create(&user).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(err.Error()))
        return
	}
}

func NewSeedController() SeedController {
	db := infrastructure.GetDB()
	return &seedController{
		db: db,
	}
}
