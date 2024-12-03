package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/db"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/product"
	"github.com/gorilla/mux"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar .env")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3306"
		fmt.Println("No se especifico el puerto, usando el 3306")
	}

	r := mux.NewRouter()

	db, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	productController := product.NewStoreDB(db)
	product.GetProductRoutes(r, productController)

	fmt.Println("Servidor iniciado en el puerto " + port)
	http.ListenAndServe(os.Getenv("DB_HOST")+":"+port, r)
}
