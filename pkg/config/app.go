package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func connect() {
	d, err := gorm.Open("mysql", "username:password@/books?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	if db == nil {
		connect()
	}
	return db
}
