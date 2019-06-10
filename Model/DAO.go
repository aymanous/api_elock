package model

import (
	apiService "../Services/API"
	"github.com/jinzhu/gorm"
)

type badge struct {
	ID        string `gorm:"primary_key"`
	FirstName string
	LastName  string
	gorm.Model
}

func (src badge) toDTO(dst *apiService.Badge) {
	dst.ID = src.ID
	dst.FirstName = src.FirstName
	dst.LastName = src.LastName
	dst.CreatedAt = src.CreatedAt
}

type settings struct {
	Server string `gorm:"primary_key"`
	Mode   string
}

type logs struct {
	gorm.Model
	Code    string
	Mode    string
	Success bool
}
