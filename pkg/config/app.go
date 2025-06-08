package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {

	d, err := gorm.Open("mysql", "root:@tcp(db:3306)/go_mysql?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}
	return db
}
