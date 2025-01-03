package order

import (
	"log"
	"net/http"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/product"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
)

type Handler struct {
	store types.OrderStore
	product types.ProductStore
}

func NewHandler(store types.OrderStore, product types.ProductStore) *Handler {
	return &Handler{store: store, product: product}
}

func (h *Handler) HandleGetOrdersUndoneRoutes(w http.ResponseWriter, r *http.Request) {
	orders, err := h.store.GetOrdersUndone()
	if err != nil {
		log.Println(err)
		return
	}
	
}