package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gihub.com/gorilla/mux"
	"github.com/kushal88053/Go_PROJECT_2/pkg/models"
	"github.com/kushal88053/Go_PROJECT_2/pkg/utils"
)

var NewBook models.Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func getBooksById(w http.ResponseWriter, r *http.Request) {
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

	models.DeleteBook(Id)
	w.WriteHeader(http.StatusNoContent)
}
