package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func initDSN() (string, error){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar .env")
		return "", err
	}
	var dsn = "host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME")
	return dsn, nil
}

func Connection() (*gorm.DB, error) {
	dsn, err := initDSN()
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	
	return db, nil
}