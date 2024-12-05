package filters

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
	"gorm.io/gorm"
)

type StoreDB struct {
	db *gorm.DB
}

func NewStoreDB(db *gorm.DB) *StoreDB {
	return &StoreDB{db: db}
}

// Types methods //
func (DB *StoreDB) HandleGetTypes(w http.ResponseWriter, r *http.Request) {
	var types []types.Type

	DB.db.Find(&types)

	err := json.NewEncoder(w).Encode(types)
	if err != nil {
		log.Println(err)
	}

}

func (DB *StoreDB) HandleCreateType(w http.ResponseWriter, r *http.Request) {
	var type_ types.Type

	err := json.NewDecoder(r.Body).Decode(&type_)
	if err != nil {
		log.Println(err)
	}
	
	DB.db.Create(&type_)
}


// ArtWork methods //
func (DB *StoreDB) HandleGetArtWorks(w http.ResponseWriter, r *http.Request) {
	var artworks []types.ArtWork

	DB.db.Find(&artworks)

	err := json.NewEncoder(w).Encode(artworks)
	if err != nil {
		log.Println(err)
	}

}

func (DB *StoreDB) HandleCreateArtWork(w http.ResponseWriter, r *http.Request) {
	var artwork types.ArtWork

	err := json.NewDecoder(r.Body).Decode(&artwork)
	if err != nil {
		log.Println(err)
	}
	
	DB.db.Create(&artwork)
}

// Category methods //
func (DB *StoreDB) HandleGetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []types.Category

	DB.db.Find(&categories)

	err := json.NewEncoder(w).Encode(categories)
	if err != nil {
		log.Println(err)
	}

}

func (DB *StoreDB) HandleCreateCategory(w http.ResponseWriter, r *http.Request) {
	var category types.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Println(err)
	}

	DB.db.Create(&category)
}