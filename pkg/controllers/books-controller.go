package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kushal88053/Go_PROJECT_2/pkg/models"
	"github.com/kushal88053/Go_PROJECT_2/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooksById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookId := vars["id"]
	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing book id", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, _ := models.GetBookById(Id)

	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	CreateBook := models.Book{}
	// Parse the request body into the NewBook struct
	if err := utils.ParseBody(r, &CreateBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book := models.CreateBook(&CreateBook)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookId := vars["id"]
	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing book id", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, err := models.DeleteBook(Id)

	if err != nil {
		fmt.Printf("error while deleting book with ID %d: %v", Id, err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Book with ID %d not found", Id)))
		return
	}
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	var updatedBook = &models.Book{}
	utils.ParseBody(r, updatedBook)
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing ...")
	}

	book, db := models.GetBookById(ID)
	if book == nil {
		fmt.Printf("book with ID %d not found", ID) // book not found
	}
	if updatedBook.Name != "" {
		book.Name = updatedBook.Name // update name if provided
	}

	if updatedBook.Author != "" {
		book.Author = updatedBook.Author // update name if provided
	}

	if updatedBook.Publication != "" {
		book.Publication = updatedBook.Publication
	}

	if err := db.Save(&book).Error; err != nil {
		fmt.Printf("fail to update") // book not found
	}

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
