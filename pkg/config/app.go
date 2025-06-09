package config

import (
	"fmt" // Used for string formatting, especially for the DSN
	"log" // For logging fatal errors
	"os"

	_ "github.com/go-sql-driver/mysql"

	// GORM v2 imports
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// Connect establishes the database connection
func Connect() {

	dbUser := os.Getenv("MYSQL_USER")
	if dbUser == "" {
		dbUser = "root" // Default user
	}

	dbPassword := os.Getenv("MYSQL_PASSWORD")
	if dbPassword == "" {

		dbPassword = "12345678"
	}

	dbHost := os.Getenv("MYSQL_HOST")
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}

	dbPortStr := os.Getenv("MYSQL_PORT")
	if dbPortStr == "" {
		dbPortStr = "3307" // Default MySQL port as a string
	}

	dbName := os.Getenv("MYSQL_DATABASE")
	if dbName == "" {
		dbName = "go_mysql" // Default database name, consistent with docker-compose.yml
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPortStr, dbName)

	var err error
	// Open the database connection using GORM v2 syntax
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		// Using log.Fatalf is better than panic in most application startup scenarios,
		// as it logs the error before exiting.
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Optional: Ping the database to ensure the connection is actually established
	// and to catch initial connection errors immediately.
	sqlDB, err := db.DB() // Get the generic database object from GORM
	if err != nil {
		log.Fatalf("Failed to get generic database object from GORM: %v", err)
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the database!")

	// Optional: Configure connection pool settings (good practice for performance)
	// sqlDB.SetMaxIdleConns(10)  // Maximum number of connections in the idle connection pool
	// sqlDB.SetMaxOpenConns(100) // Maximum number of open connections to the database
	// sqlDB.SetConnMaxLifetime(time.Hour) // Maximum amount of time a connection may be reused
}

// GetDB returns the established GORM database instance.
// It ensures that Connect() is called if the connection hasn't been made yet.
func GetDB() *gorm.DB {
	// This check also helps ensure that the `db` variable is not nil
	// if `Connect()` somehow failed to set it (though `log.Fatalf` should prevent that).
	if db == nil {
		Connect() // Ensure connection is established
	}
	return db
}
