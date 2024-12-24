package cart

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/auth"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/utils"
)

type Handler struct {
	store types.CartStore
}

func NewHandler(store types.CartStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) CreateCart(w http.ResponseWriter, r *http.Request) {
	var cart types.CartPayload
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		log.Println(err)
		return
	}
	if role := auth.RoleUser(cart.Token); role != "admin" {
		utils.UnauthorizedUser(w)
		return
	}

	user, err := auth.ParseToken(cart.Token)
	if err != nil {
		log.Println(err)
		return
	}
	var total float64
	for i := 0; i < len(cart.Productid); i++ {
		total += float64(cart.Quantity[i]) * cart.Price[i]
	}

	cartId, err := h.store.CreateCart(int64(user.ID), total, cart.Address)
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < len(cart.Productid); i++ {
		itemId, _ := h.store.CreateCartItem(int64(cart.Productid[i]), cart.Quantity[i], cart.Price[i])
		err = h.store.CreateOrder(cartId, itemId)
		if err != nil {
			log.Println(err)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"cart_id": %d}`, cartId)))
}
