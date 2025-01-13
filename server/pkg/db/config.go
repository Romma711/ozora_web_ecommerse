package db

import (
	"database/sql"
	"log"
	_ "os"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

///Esto es la conexion con la base de datos de turso
/*
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
*/
///Esto es la conexion con la base de datos de prueba

func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("sqlite", "./ozora_db_testing.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
