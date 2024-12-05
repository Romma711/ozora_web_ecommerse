package product

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type StoreDB struct {
	db *gorm.DB
}

func NewStoreDB(db *gorm.DB) *StoreDB {
	return &StoreDB{db: db}
}

func (DB *StoreDB) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	var products []types.Product

	DB.db.Find(&products)

	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Println(err)
	}

}

func (DB *StoreDB) HandleGetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var product types.Product
	DB.db.First(&product, params["id"])
	err := json.NewEncoder(w).Encode(product)
	if err != nil {
		log.Println(err)
	}

}

func (DB *StoreDB) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var product types.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println(err)
	}

	DB.db.Create(&product)
}

func (DB *StoreDB) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var product types.Product
	DB.db.First(&product, params["id"])

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println(err)
	}

	DB.db.Save(&product)
}

func (DB *StoreDB) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var product types.Product
	DB.db.First(&product, params["id"])

	DB.db.Delete(&product)
}

func (db *StoreDB) DB() *gorm.DB {
	return db.db
}
