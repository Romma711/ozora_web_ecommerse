package db

import (
	"database/sql"
	"log"
	"os"
	"github.com/joho/godotenv"
)


func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("TURSO_DATABASE_URL")+"?authToken="+os.Getenv("TURSO_AUTH_TOKEN")

  	db, err := sql.Open("libsql", url)
  	if err != nil {
    	return nil, err
  	}
	
	return db, nil
}