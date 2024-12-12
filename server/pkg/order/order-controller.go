package order
/*

import (
	"database/sql"
	"encoding/json"
	"go/types"
	"log"
	"net/http"
)

type StoreDB struct {
	db *sql.DB
	// db *sql.DB
}

func NewStoreDB(db *sql.DB) *StoreDB {
	return &StoreDB{db: db}
}
func (db *StoreDB) HandleGetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []types.Order

	//Funcion que devuelve todas las ordenes

	err := json.NewEncoder(w).Encode(orders)
	if err != nil {
		log.Println(err)
	}
}

func (db *StoreDB) HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	var order types.Order

	//fuincion que devuelve una orden en especifico

	err := json.NewEncoder(w).Encode(order)
	if err != nil {
		log.Println(err)
	}
}

func (db *StoreDB) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var order types.Order

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Println(err)
	}
	
	//Funcion que crea una orden
}

func (db *StoreDB) HandleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	var order types.Order

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Println(err)
	}
	
	//Funcion para borrar una orden
}
	*/