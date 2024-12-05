package filters

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetFiltersRoutes(r *mux.Router, db *StoreDB) {
	r.HandleFunc("/product/categories", db.HandleGetCategories).Methods(http.MethodGet)
	r.HandleFunc("/product/categories", db.HandleCreateCategory).Methods(http.MethodPost)
	r.HandleFunc("/product/types", db.HandleGetTypes).Methods(http.MethodGet)
	r.HandleFunc("/product/types", db.HandleCreateType).Methods(http.MethodPost)
	r.HandleFunc("/product/artworks", db.HandleGetArtWorks).Methods(http.MethodGet)
	r.HandleFunc("/product/artworks", db.HandleCreateArtWork).Methods(http.MethodPost)
}