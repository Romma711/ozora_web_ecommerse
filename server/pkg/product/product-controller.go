package product

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) GetProductRoutes(r *mux.Router) {
	r.HandleFunc("/products", h.HandleGetProducts).Methods(http.MethodGet)
	r.HandleFunc("/products/{id}", h.HandleGetProduct).Methods(http.MethodGet)
	r.HandleFunc("/products", h.HandleCreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/products/{id}", h.HandleUpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/products/{id}", h.HandleDeleteProduct).Methods(http.MethodDelete)
}

func (h *Handler) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	var products []types.Product

	//Funcion que devuelve todos los productos

	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Println(err)
	}

}

func (h *Handler) HandleGetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var product types.Product
	//Funcion que devuelve un producto en especifico
	err := json.NewEncoder(w).Encode(product)
	if err != nil {
		log.Println(err)
	}

}

func (h *Handler) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var product types.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println(err)
	}

	//Funcion que crea un producto
}

func (h *Handler) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var product types.Product
	//Funcion que actualiza un producto

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println(err)
	}

}

func (h *Handler) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var product types.Product
	//Funcion para borrar un producto

}

