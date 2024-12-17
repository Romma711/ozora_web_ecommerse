package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/db"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/filters"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/product"
	"github.com/gorilla/mux"
)

func main() {
	
	db, err :=db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	
	r := mux.NewRouter()

	productStore := product.NewStoreDB(db)
	filtersStore := filters.NewStore(db)

	productHandler := product.NewHandler(productStore)
	filtersHandler := filters.NewHandler(filtersStore)


	productHandler.GetProductRoutes(r)
	filtersHandler.GetFiltersRoutes(r)
	fmt.Println("Servidor iniciado en el puerto " + os.Getenv("DB_PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT"), r))
}
