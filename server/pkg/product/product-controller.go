package product

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/auth"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) GetProductRoutes(r *mux.Router) {
	r.HandleFunc("/products", h.HandleGetProducts).Methods(http.MethodGet)
	r.HandleFunc("/product/{id}", h.HandleGetProduct).Methods(http.MethodGet)
	r.HandleFunc("/products/tag/{id}", h.HandleGetProductsFiltered).Methods(http.MethodGet)

	///ADMIN AND EMPLOYEES ROUTES
	r.HandleFunc("/products", h.HandleCreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/products/{id}", h.HandleUpdateProduct).Methods(http.MethodPut)
}

func (h *Handler) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	var products []types.Product

	products, err := h.store.GetProducts()
	if err != nil {
		log.Println(err)
	}

	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Println(err)
	}

}

func (h *Handler) HandleGetProductsFiltered(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var products []types.Product
	var err error
	id, _ := strconv.Atoi(params["id"])
	filter := params["filter"]
	if filter == "category" {
		products, _ = h.store.GetProductsByCategory(id)
	}
	if filter == "type" {
		products, _ = h.store.GetProductsByTypes(id)
	}
	if filter == "artwork" {
		products, _ = h.store.GetProductsByArtWork(id)
	}
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Println(err)
	}

}

func (h *Handler) HandleGetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	if id == 0 {
		return
	}
	product, err := h.store.GetProductByID(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("El producto que busca no existe")
		return
	}
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	token := mux.Vars(r)["token"]
	if role := auth.RoleUser(token); role != "admin" && role != "employee" {
		utils.UnauthorizedUser(w)
		return
	}
	
	var product types.ProductPayLoad
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println(err)
	}

	err = h.store.CreateProduct(&product)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"Product created successfully\"}"))
}

func (h *Handler) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	token := mux.Vars(r)["token"]
	if role := auth.RoleUser(token); role != "admin" && role != "employee" {
		utils.UnauthorizedUser(w)
		return
	}

	var product types.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println(err)
	}
	err = h.store.UpdateProduct(&product)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"Product updated successfully\"}"))
}
