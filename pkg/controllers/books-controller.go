package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kushal88053/Go_PROJECT_2/pkg/models" // Assuming this is your models package
	"github.com/kushal88053/Go_PROJECT_2/pkg/utils"  // Assuming this package handles ParseBody
)

// Helper function for consistent error responses
func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// Helper function for consistent success responses
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

// GetBooks fetches all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks() // Now returns an error
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve books from database")
		return
	}
	respondWithJSON(w, http.StatusOK, books)
}

// GetBooksById fetches a single book by ID
func GetBooksById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookIDStr := vars["id"]
	bookID, err := strconv.ParseInt(bookIDStr, 10, 64) // Base 10, 64-bit integer

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID provided")
		return
	}

	book, err := models.GetBookById(bookID) // Now returns an error
	if err != nil {
		// Distinguish between "not found" and other database errors if needed
		// For GORM, you might check if err == gorm.ErrRecordNotFound
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving book: %v", err))
		return
	}
	if book == nil { // Book not found
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("Book with ID %d not found", bookID))
		return
	}

	respondWithJSON(w, http.StatusOK, book)
}

// CreateBook creates a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var createBook models.Book // Renamed variable for clarity
	if err := utils.ParseBody(r, &createBook); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	// models.CreateBook now handles potential errors internally (logs them),
	// but in a real-world scenario, you might want it to return an error
	// to allow the controller to respond accordingly. For now, we rely on its internal logging.
	book := models.CreateBook(&createBook)

	// Assuming CreateBook always returns a *models.Book instance
	// and logging any internal errors. If it were to return (book *Book, err error),
	// you'd check err here.

	respondWithJSON(w, http.StatusCreated, book)
}

// DeleteBook deletes a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookIDStr := vars["id"]
	bookID, err := strconv.ParseInt(bookIDStr, 10, 64)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID provided")
		return
	}

	deletedBook, err := models.DeleteBook(bookID) // Now returns an error
	if err != nil {
		// The models.DeleteBook now returns a more specific error for "not found"
		respondWithError(w, http.StatusNotFound, err.Error()) // Use the error message from models
		return
	}

	// If deletion was successful, return a success message or the deleted book
	respondWithJSON(w, http.StatusOK, map[string]string{"message": fmt.Sprintf("Book with ID %d deleted successfully", bookID)})
	// Or, if you want to return the deleted book object:
	respondWithJSON(w, http.StatusOK, deletedBook)
}

// UpdateBook updates an existing book by ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookIDStr := vars["id"]
	bookID, err := strconv.ParseInt(bookIDStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID provided for update")
		return
	}

	// Fetch the existing book first
	bookToUpdate, err := models.GetBookById(bookID)
	if err != nil {
		// This handles database errors during retrieval, not "not found"
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving book for update: %v", err))
		return
	}
	if bookToUpdate == nil { // Book not found
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("Book with ID %d not found for update", bookID))
		return
	}

	// Parse the request body for update fields
	var requestBody struct { // Using an anonymous struct for parsing partial updates
		Name        string `json:"name"`
		Author      string `json:"author"`
		Publication string `json:"publication"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	// Apply updates only if the fields are provided in the request body
	// This allows partial updates
	if requestBody.Name != "" {
		bookToUpdate.Name = requestBody.Name
	}
	if requestBody.Author != "" {
		bookToUpdate.Author = requestBody.Author
	}
	if requestBody.Publication != "" {
		bookToUpdate.Publication = requestBody.Publication
	}

	// Call the new UpdateBookInDB function from the models package
	if err := models.UpdateBookInDB(bookToUpdate); err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to update book in database: %v", err))
		return
	}

	// Respond with the updated book object
	respondWithJSON(w, http.StatusOK, bookToUpdate)
}
