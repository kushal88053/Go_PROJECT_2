package models

import (
	"github.com/jinzhu/gorm"
	"github.com/kushal88053/Go_PROJECT_2/pkg/config"
)

var (
	db *gorm.DB
)

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:autor`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})

}

func CreateBook(book *Book) *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}
func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db.Where("ID=?", Id).Find(&book)
	if book.ID != 0 {
		return &book, db
	}
	return nil, db

}

func DeleteBook(Id int64) {
	var book Book
	db.Where("ID=?", Id).Delete(&book)

}
