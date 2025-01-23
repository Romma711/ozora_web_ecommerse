package filters

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	r.HandleFunc("/recomendation/{random}", h.HandleGetRecomandation).Methods(http.MethodGet)
	r.HandleFunc("/product/artworks", h.HandleGetArtWorks).Methods(http.MethodGet)
	r.HandleFunc("/product/categories", h.HandleGetCategories).Methods(http.MethodGet)
	r.HandleFunc("/product/types", h.HandleGetTypes).Methods(http.MethodGet)

	//Admin routes
	r.HandleFunc("/admin/product/categories", h.HandleCreateCategory).Methods(http.MethodPost)
	r.HandleFunc("/admin/product/categories", h.HandleUpdateCategory).Methods(http.MethodPut)
	r.HandleFunc("/admin/product/types", h.HandleCreateType).Methods(http.MethodPost)
	r.HandleFunc("/admin/product/types", h.HandleUpdateType).Methods(http.MethodPut)
	r.HandleFunc("/admin/product/artworks", h.HandleCreateArtWork).Methods(http.MethodPost)
	r.HandleFunc("/admin/product/artworks", h.HandleUpdateArtWork).Methods(http.MethodPut)
}

// Types methods //

//Funcion que muestra todos los types
func (h *Handler) HandleGetTypes(w http.ResponseWriter, r *http.Request) {
	types, err := h.store.GetTypes()
	if err != nil {
		log.Println(err)
		return
	}

	err = json.NewEncoder(w).Encode(types)
	if err != nil {
		log.Println(err)
	}
}

//Funcion que crea un tipo de producto
func (h *Handler) HandleCreateType(w http.ResponseWriter, r *http.Request) {
	var type_ types.Type

	err := json.NewDecoder(r.Body).Decode(&type_)
	if err != nil {
		log.Println(err)
	}
	
	err = h.store.CreateType(&type_)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"Type created successfully\"}"))
}

func (h *Handler) HandleUpdateType(w http.ResponseWriter, r *http.Request){
	var type_ types.Type
	err :=json.NewDecoder(r.Body).Decode(&type_)
	if err != nil{
		log.Println(err)
	}

	err = h.store.UpdateType(type_)
	if err != nil{
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"Type updated successfully\"}"))

}

// ArtWork methods //
func (h *Handler) HandleGetArtWorks(w http.ResponseWriter, r *http.Request) {
	artworks, err := h.store.GetArtWorks()
	if err != nil {
		log.Println(err)
		return
	}

	err = json.NewEncoder(w).Encode(artworks)
	if err != nil {
		log.Println(err)
	}

}
//This function return an artwork as recomendation
func (h *Handler) HandleGetRecomandation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number ,_ := strconv.Atoi(vars["random"])
	artwork, err := h.store.GetArtWorkRecomendation(number)
	if err != nil {
		log.Println(err)
		return
	}

	if artwork.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\":\"No artwork found\"}"))
		return
	}
	
	err = json.NewEncoder(w).Encode(artwork)
	if err != nil {
		log.Println(err)
	}
}

//Funcion que crea un artwork
func (h *Handler) HandleCreateArtWork(w http.ResponseWriter, r *http.Request) {
	var artwork types.ArtWork

	err := json.NewDecoder(r.Body).Decode(&artwork)
	if err != nil {
		log.Println(err)
	}
	
	err = h.store.CreateArtWork(&artwork)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")	
	w.Write([]byte("{\"message\":\"Artwork created successfully\"}"))
}

func (h *Handler) HandleUpdateArtWork(w http.ResponseWriter, r *http.Request){
	var artwork types.ArtWork
	err :=json.NewDecoder(r.Body).Decode(&artwork)
	if err != nil{
		log.Println(err)
	}

	err = h.store.UpdateArtWork(artwork)
	if err != nil{
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"Artwork updated successfully\"}"))

}

// Category methods //
func (h *Handler) HandleGetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.store.GetCategories()
	if err != nil {
		log.Println(err)
		return
	}

	err = json.NewEncoder(w).Encode(categories)
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

	err = h.store.CreateCategory(&category)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"Category created successfully\"}"))
}

func (h *Handler) HandleUpdateCategory(w http.ResponseWriter, r *http.Request){
	var category types.Category
	err :=json.NewDecoder(r.Body).Decode(&category)
	if err != nil{
		log.Println(err)
	}

	err = h.store.UpdateCategory(category)
	if err != nil{
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"Category updated successfully\"}"))
}