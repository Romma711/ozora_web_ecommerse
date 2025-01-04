package order

import (
	"database/sql"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}


func (s *Store) GetOrdersUndone() ([]types.Cart, error) {
	rows, err := s.db.Query(`SELECT id FROM cart WHERE status = 'IN PROCESS'`)
	if err != nil {
		return nil, err
	}

	carts := make([]types.Cart, 0)

	for rows.Next() {
		cart, err := ScanRowsIntoCart(rows)
		if err != nil {
			return nil, err
		}
		carts = append(carts, *cart)
	}

	return carts, nil
}

func (s *Store) GetOrderByOrderId(cartId int) ([]types.Order, error) {
	rows, err := s.db.Query(`SELECT id_product, quantity, price  
							 FROM cart_shopping  
							 WHERE id_cart = ?`, cartId)
	if err != nil {
		return nil, err
	}

	orders := make([]types.Order, 0)

	for rows.Next() {
		order, err := ScanRowsIntoOrder(rows)
		if err != nil {
			return nil, err
		}
		orders = append(orders, *order)
	}
	return orders, nil
}

func (s *Store) GetOrdersByUserId(userId int) ([]types.Order, error) {
	rows, err := s.db.Query(`SELECT id_product, quantity, price, id_cart  
							 FROM cart_shopping 
							 WHERE id_cart IN (SELECT id FROM cart WHERE id_client = ?)`, userId)
	if err != nil {
		return nil, err
	}

	orders := make([]types.Order, 0)

	for rows.Next() {
		order, err := ScanRowsIntoOrder(rows)
		if err != nil {
			return nil, err
		}
		orders = append(orders, *order)
	}
	return orders, nil
}

func ScanRowsIntoOrder(row *sql.Rows) (*types.Order, error) {
	order := new(types.Order)
	err := row.Scan(
		&order.IDProduct,
		&order.Price,
		&order.Quantity,
		&order.IDCart,
	)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func ScanRowsIntoCart(row *sql.Rows) (*types.Cart, error) {
	cart := new(types.Cart)
	err := row.Scan(
		&cart.ID,
		&cart.UserID,
		&cart.Total,
		&cart.Address,
		&cart.Status,
	)
	if err != nil {
		return nil, err
	}
	return cart, nil
}