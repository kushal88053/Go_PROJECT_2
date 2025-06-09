package models

import (
	"fmt"
	"log" // For logging if needed, or you can just return errors

	"github.com/kushal88053/Go_PROJECT_2/pkg/config"
	"gorm.io/gorm" // <--- UPDATED: This is the GORM v2 import
)

var (
	db *gorm.DB // This will hold our GORM database instance
)

// Book struct represents the book model in your database
type Book struct {
	gorm.Model         // GORM's embedded model for ID, CreatedAt, UpdatedAt, DeletedAt
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	log.Println("Initializing models package...")

	db = config.GetDB()

	if db == nil {
		log.Fatalf("Database connection (db instance) is nil before AutoMigrate. This indicates a problem in config.Connect() or GetDB().")
	}

	// --- ADD THIS EXACT LINE ---
	fmt.Printf("DEBUG: Before AutoMigrate, db instance is %v (type %T)\n", db, db)
	// --- END OF ADDED LINE ---

	db.AutoMigrate(&Book{}) // This is line 37 (or close)
	// if err != nil {                      // This is line 38 (or close)
	// 	log.Fatalf("Failed to auto migrate database for Book model: %v", err)
	// }

	log.Println("Database migration for Book model completed.")
}

// CreateBook inserts a new book record into the database
func CreateBook(book *Book) *Book {
	// In GORM v2, db.Create handles setting primary keys and timestamps (from gorm.Model)
	// You no longer need db.NewRecord(book)
	result := db.Create(book)
	if result.Error != nil {
		log.Printf("Error creating book: %v", result.Error)
		// Depending on your application, you might want to return an error here
		// or handle it differently. For simplicity, we'll just log.
	}
	return book // The 'book' pointer will now contain the updated ID and timestamps
}

// GetAllBooks fetches all book records from the database
func GetAllBooks() ([]Book, error) { // Added error return
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		log.Printf("Error getting all books: %v", result.Error)
		return nil, result.Error // Return nil and the error
	}
	return books, nil // Return books and no error
}

// GetBookById fetches a single book record by its ID
// Returns the book pointer and an error (nil if found, gorm.ErrRecordNotFound if not found, or other error)
func GetBookById(Id int64) (*Book, error) { // Updated signature to return error
	var book Book
	result := db.Where("ID = ?", Id).First(&book) // Use First for single record by primary key

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Printf("Error getting book by ID %d: %v", Id, result.Error)
		return nil, result.Error // Return nil book and the actual database error
	}

	return &book, nil
}

// DeleteBook deletes a book record by its ID
// Returns the deleted book pointer and an error
func DeleteBook(id int64) (*Book, error) {

	var book Book

	// First, fetch the book to return it after deletion (optional, but requested)
	result := db.Where("id = ?", id).First(&book)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("book with ID %d not found", id) // More user-friendly error
		}
		log.Printf("Error finding book to delete with ID %d: %v", id, result.Error)
		return nil, result.Error // Return the actual database error
	}

	// Then, delete it
	deleteResult := db.Delete(&book)
	if deleteResult.Error != nil {
		log.Printf("Error deleting book with ID %d: %v", id, deleteResult.Error)
		return nil, deleteResult.Error // Return the error if deletion fails
	}
	// You can also check deleteResult.RowsAffected > 0 to confirm deletion happened

	return &book, nil // Return the deleted book and no error
}

func UpdateBookInDB(book *Book) error {
	result := db.Save(book) // db here refers to the package-level db in models.go
	return result.Error     // Return any error that occurred during the save operation
}
