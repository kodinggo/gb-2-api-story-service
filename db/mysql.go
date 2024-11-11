package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kodinggo/gb-2-api-story-service/internal/helper"
)

// NewMysql is a function to initialize the MySQL database.
func NewMysql() *sql.DB {
	// Initialize the database
	db, err := sql.Open("mysql", helper.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	// Check database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
