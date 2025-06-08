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
	Name        string `json:"name"`
	Author      string `json:"author"`
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

func DeleteBook(id int64) (*Book, error) {
	var book Book
	// First, fetch the book
	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err // book not found or some DB error
	}

	// Then, delete it
	if err := db.Delete(&book).Error; err != nil {
		return nil, err // failed to delete
	}

	return &book, nil // return the deleted book
}
