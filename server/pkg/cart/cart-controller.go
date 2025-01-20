package cart

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/auth"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
	//"github.com/Romma711/ozora_web_ecommerse/server/pkg/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.CartStore
	product types.ProductStore
}

func NewHandler(store types.CartStore, product types.ProductStore) *Handler {
	return &Handler{store: store, product: product}
}

func (h *Handler) GetCartRoutes(r *mux.Router) {
	r.HandleFunc("/cart", h.HandleCreateCart).Methods(http.MethodPost)
}

func (h *Handler) HandleCreateCart(w http.ResponseWriter, r *http.Request) {
	var cart types.CartPayload
	err := json.NewDecoder(r.Body).Decode(&cart) //decodifica el json que viene en el body
	if err != nil {
		log.Println(err)
		return
	}
	/*if role := auth.RoleUser(cart.Token); role != "client" {
		utils.UnauthorizedUser(w)
		return
	}
	*/

	user, err := auth.ParseToken(cart.Token) //parsea el token
	if err != nil {
		log.Println(err)
		return
	}
	var total float64
	for i := 0; i < len(cart.ProductsCart); i++ {	//recorre los productos del carrito
		product, err := h.product.GetProductByID(cart.ProductsCart[i].ID) //obtiene el producto por id
		if err != nil {
			log.Println(err)
			return
		}
		cart.ProductsCart[i].Total = float64(cart.ProductsCart[i].Quantity) * product.Price //calcula el total del producto
		total += cart.ProductsCart[i].Total //suma el total del producto al total del carrito
	}

	cartId, err := h.store.CreateCart(user.ID, total, cart.Address) //crea el carrito
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < len(cart.ProductsCart); i++ { //recorre los productos del carrito
		err := h.store.CreateCartItem(cart.ProductsCart[i].ID, cart.ProductsCart[i].Quantity, cart.ProductsCart[i].Total, cartId) //crea el item del carrito
		if err != nil {
			log.Println(err)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"cart_id": %d}`, cartId)))
}
