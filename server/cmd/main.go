package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/db"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/filters"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/product"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

	db.AutoMigrate(&types.Product{}, &types.Category{}, &types.Type{}, &types.ArtWork{})

	productDB := product.NewStoreDB(db)
	filtersDB := filters.NewStoreDB(db)
	product.GetProductRoutes(r, productDB)
	filters.GetFiltersRoutes(r, filtersDB)

	fmt.Println("Servidor iniciado en el puerto " + port)
	http.ListenAndServe(os.Getenv("DB_HOST")+":"+port, r)
}
