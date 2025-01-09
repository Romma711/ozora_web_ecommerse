package order

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	store   types.OrderStore
	product types.ProductStore
}

func NewHandler(store types.OrderStore, product types.ProductStore) *Handler {
	return &Handler{store: store, product: product}
}

func (h *Handler) GetRoutes(mux *mux.Router) {
	mux.HandleFunc("/admin/orders", h.HandleGetOrdersUndoneRoutes).Methods("GET")
	mux.HandleFunc("/admin/orders/:id", h.HandleGetOrderByOrderId).Methods("GET")
	mux.HandleFunc("/admin/orders/user/:id", h.HandleGetOrdersByUserId).Methods("GET")
}

func (h *Handler) HandleGetOrdersUndoneRoutes(w http.ResponseWriter, r *http.Request) {
	orders, err := h.store.GetOrdersUndone()
	if err != nil {
		log.Println(err)
		return
	}

	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandleGetOrderByOrderId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	orders, err := h.store.GetOrderByOrderId(id)
	if err != nil {
		log.Println(err)
		return
	}
	var response types.OrderResponse
	for _, order := range orders {
		product, err := h.product.GetProductByID(order.IDProduct)
		if err != nil {
			log.Println(err)
			return
		}
		response.Products = append(response.Products, *product)
		response.Quantity = append(response.Quantity, order.Quantity)
		response.Total += order.Price * float64(order.Quantity)
	}
	response.IDCart = orders[0].IDCart

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandleGetOrdersByUserId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	orders, err := h.store.GetOrdersByUserId(id)
	if err != nil {
		log.Println(err)
		return
	}

	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}