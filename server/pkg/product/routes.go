package product

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetProductRoutes(r *mux.Router, db *StoreDB) {
	r.HandleFunc("/products", db.HandleGetProducts).Methods(http.MethodGet)
	r.HandleFunc("/products/{id}", db.HandleGetProduct).Methods(http.MethodGet)
	r.HandleFunc("/products", db.HandleCreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/products/{id}", db.HandleUpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/products/{id}", db.HandleDeleteProduct).Methods(http.MethodDelete)
}

