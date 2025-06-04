package models

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

type Book struct {
	gorm.Model
	Name   string `gorm:""json:"name"`
	Author string `json:`
}
