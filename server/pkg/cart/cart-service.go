package cart

import (
	"database/sql"
)

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{db: db}
}

func (s *store) CreateCartItem(productId int, quantity int, price float64, cartId int)  error {
	_, err := s.db.Exec(`INSERT INTO cart_shopping (product_id, quantity, price, id_cart) VALUES (?, ?, ?, ?)`, productId, quantity, price, cartId)
	if err != nil {
		return err
	}
	return nil
}

func (s *store) CreateCart(userId int, total float64, address string) (int, error) {
	res, err := s.db.Exec(`INSERT INTO cart (id_client, address) VALUES ($1, $2, $3, $4)`, userId, total, address)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), err
}
