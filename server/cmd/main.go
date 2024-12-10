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
)

func main() {
	
	db, err :=db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	
	r := mux.NewRouter()

	productDB := product.NewStoreDB(db)
	filtersDB := filters.NewStoreDB(db)

	product.GetProductRoutes(r, productDB)
	filters.GetFiltersRoutes(r, filtersDB)

	fmt.Println("Servidor iniciado en el puerto " + os.Getenv("DB_PORT"))
	http.ListenAndServe(os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT"), r)
}
