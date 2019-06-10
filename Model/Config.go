package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type DBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	Driver   string `json:"driver"`
}
type Model struct {
	db *gorm.DB
}

func NewModel(config DBConfig) (model *Model, err error) {
	model = &Model{}
	connect := fmt.Sprintf(
		"%v:%v@/%v?charset=utf8&parseTime=True&loc=Local",
		config.Username, config.Password, config.Name)
	model.db, err = gorm.Open(config.Driver, connect)
	if err != nil {
		return nil, err
	}

	tdb := model.db.Begin()
	defer tdb.Rollback()

	if err = tdb.AutoMigrate(
		&badge{},
		&settings{},
		&logs{},
	).Error; err != nil {
		return nil, err
	}

	tdb.Commit()
	return model, nil
}
