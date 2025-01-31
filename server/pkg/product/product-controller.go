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
	tag   types.TagsStore
}

func NewHandler(store types.ProductStore, tag types.TagsStore) *Handler {
	return &Handler{store: store, tag: tag}
}

func (h *Handler) GetProductRoutes(r *mux.Router) {
	r.HandleFunc("/products", h.HandleGetProducts).Methods(http.MethodGet)
	r.HandleFunc("/products/{id}", h.HandleGetProduct).Methods(http.MethodGet)
	r.HandleFunc("/products/tag", h.HandleGetProductsFiltered).Methods(http.MethodGet)
	r.HandleFunc("/products/newer/", h.HandleGetProductDatetime).Methods(http.MethodGet)
	r.HandleFunc("/products/noted/", h.HandleGetProductSold).Methods(http.MethodGet)

	///ADMIN AND EMPLOYEES ROUTES
	r.HandleFunc("/admin/products/create", h.HandleCreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/admin/products/change/{id}", h.HandleUpdateProduct).Methods(http.MethodPut)
}

func (h *Handler) HandleGetProductDatetime(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProductsByDatetime()
	if err != nil {
		log.Println("Error:" + err.Error())
		return
	}

	var productsResponse []types.ProductResponse

	for i := 0; i < len(products); i++ {
		response, err := h.ReturnProduct(products[i])
		if err != nil {
			log.Println("Error:" + err.Error())
			return
		}
		productsResponse = append(productsResponse, response)
	}

	_ = json.NewEncoder(w).Encode(&productsResponse)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (h *Handler) HandleGetProductSold(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProductsBySold()
	if err != nil {
		log.Println("Error:" + err.Error())
		return
	}

	var productsResponse []types.ProductResponse

	for i := 0; i < len(products); i++ {
		response, err := h.ReturnProduct(products[i])
		if err != nil {
			log.Println("Error:" + err.Error())
			return
		}
		productsResponse = append(productsResponse, response)
	}

	_ = json.NewEncoder(w).Encode(&productsResponse)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (h *Handler) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	var products []types.Product
	var productsResponse []types.ProductResponse

	products, err := h.store.GetProducts()
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < len(products); i++ {
		response, err := h.ReturnProduct(products[i])
		if err != nil {
			log.Println(err)
		}
		productsResponse = append(productsResponse, response)
	}

	err = json.NewEncoder(w).Encode(&productsResponse)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

// /This function is used to return a product with the tags converted as strings
func (h *Handler) ReturnProduct(product types.Product) (types.ProductResponse, error) {
	var productResponse types.ProductResponse
	var err error
	///Assigning the values
	productResponse.ID = product.ID
	productResponse.Name = product.Name
	productResponse.Description = product.Description
	productResponse.Price = product.Price
	productResponse.Stock = product.Stock
	productResponse.Image = product.Image

	///Error handling
	productResponse.Category, err = h.tag.GetCategoryById(product.CategoryID)
	if err != nil {
		log.Println(err)
		return productResponse, err
	}
	productResponse.Type, err = h.tag.GetTypeById(product.TypeID)
	if err != nil {
		return productResponse, err
	}
	productResponse.ArtWork, err = h.tag.GetArtWorkById(product.ArtWorkID)
	if err != nil {
		return productResponse, err
	}

	return productResponse, nil
}

func (h *Handler) HandleGetProductsFiltered(w http.ResponseWriter, r *http.Request) {
	var products []types.Product
	var err error
	tag := r.URL.Query().Get("tag")
	filter := r.URL.Query().Get("filter")
	
	if tag == "" || filter == "" {
		http.Error(w, "Missing query parameters", http.StatusBadRequest)
		return
	}
	log.Println(tag, filter)

	if tag == "category" {
		products, _ = h.store.GetProductsByCategory(filter)
	}
	if tag == "type" {
		products, _ = h.store.GetProductsByTypes(filter)
	}
	if tag == "artwork" {
		products, _ = h.store.GetProductsByArtWork(filter)
	}

	var productsResponse []types.ProductResponse

	for i := 0; i < len(products); i++ {
		response, err := h.ReturnProduct(products[i])
		if err != nil {
			log.Println(err)
		}
		productsResponse = append(productsResponse, response)
	}

	err = json.NewEncoder(w).Encode(&productsResponse)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

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

	productResponse, _ := h.ReturnProduct(*product)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(productResponse)
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	/*token := mux.Vars(r)["token"]
	if role := auth.RoleUser(token); role != "admin" && role != "employee" {
		utils.UnauthorizedUser(w)
		return
	}
	*/
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
