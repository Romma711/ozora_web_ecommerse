package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/cart"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/db"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/filters"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/product"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/user"
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
	userStore := user.NewStore(db)
	cartStore := cart.NewStore(db)

	productHandler := product.NewHandler(productStore)
	filtersHandler := filters.NewHandler(filtersStore)
	userHandler := user.NewHandler(userStore)
	cartHandler := cart.NewHandler(cartStore)


	productHandler.GetProductRoutes(r)
	filtersHandler.GetFiltersRoutes(r)
	userHandler.GetUsersRoutes(r)
	cartHandler.GetCartRoutes(r)
	fmt.Println("Servidor iniciado en el puerto " + os.Getenv("DB_PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT"), r))
}
