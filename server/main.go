package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/cart"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/db"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/filters"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/order"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/product"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/user"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/utils"
	"github.com/gorilla/mux"
)

func main() {

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	router := mux.NewRouter()
	subRouter := router.Host(host).PathPrefix("/api").Subrouter()
	
	productStore := product.NewStoreDB(db)
	filtersStore := filters.NewStore(db)
	userStore := user.NewStore(db)
	cartStore := cart.NewStore(db)
	orderStore := order.NewStore(db)
	
	filtersHandler := filters.NewHandler(filtersStore)
	productHandler := product.NewHandler(productStore, filtersStore)
	userHandler := user.NewHandler(userStore)
	cartHandler := cart.NewHandler(cartStore, productStore)
	orderHandler := order.NewHandler(orderStore, productStore)
	
	
	productHandler.GetProductRoutes(subRouter)
	filtersHandler.GetFiltersRoutes(subRouter)
	userHandler.GetUsersRoutes(subRouter)
	cartHandler.GetCartRoutes(subRouter)
	orderHandler.GetRoutes(subRouter)
	
	c := utils.EnableCORS()
	muxCors := c.Handler(subRouter)

	fmt.Println("Servidor iniciado en el host " + host +":"+ port)
	log.Fatal(http.ListenAndServe((host +":"+ port), muxCors))
}
