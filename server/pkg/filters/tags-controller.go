package filters

import (
	
	"encoding/json"
	"log"
	"net/http"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.TagsStore
}

func NewHandler(store types.TagsStore) *Handler {
	return &Handler{store: store}
}


func (h *Handler)GetFiltersRoutes(r *mux.Router) {
	r.HandleFunc("/product/categories", h.HandleGetCategories).Methods(http.MethodGet)
	r.HandleFunc("/product/categories", h.HandleCreateCategory).Methods(http.MethodPost)
	r.HandleFunc("/product/types", h.HandleGetTypes).Methods(http.MethodGet)
	r.HandleFunc("/product/types", h.HandleCreateType).Methods(http.MethodPost)
	r.HandleFunc("/product/artworks", h.HandleGetArtWorks).Methods(http.MethodGet)
	r.HandleFunc("/product/artworks", h.HandleCreateArtWork).Methods(http.MethodPost)
}

// Types methods //
func (h *Handler) HandleGetTypes(w http.ResponseWriter, r *http.Request) {
	var types []types.Type

	//Funcion que devuelve todos los tipos de productos

	err := json.NewEncoder(w).Encode(types)
	if err != nil {
		log.Println(err)
	}

}

func (h *Handler) HandleCreateType(w http.ResponseWriter, r *http.Request) {
	var type_ types.Type

	err := json.NewDecoder(r.Body).Decode(&type_)
	if err != nil {
		log.Println(err)
	}
	
	//Funcion que crea un tipo de producto
}


// ArtWork methods //
func (h *Handler) HandleGetArtWorks(w http.ResponseWriter, r *http.Request) {
	var artworks []types.ArtWork

	//Funcion que devuelve todos los artworks

	err := json.NewEncoder(w).Encode(artworks)
	if err != nil {
		log.Println(err)
	}

}

func (h *Handler) HandleCreateArtWork(w http.ResponseWriter, r *http.Request) {
	var artwork types.ArtWork

	err := json.NewDecoder(r.Body).Decode(&artwork)
	if err != nil {
		log.Println(err)
	}
	
	//Funcion que crea un artwork
}

// Category methods //
func (h *Handler) HandleGetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []types.Category

	//Funcion que devuelve todas las categorías

	err := json.NewEncoder(w).Encode(categories)
	if err != nil {
		log.Println(err)
	}

}

func (h *Handler) HandleCreateCategory(w http.ResponseWriter, r *http.Request) {
	var category types.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Println(err)
	}

	//Funcion que crea una categoría
}