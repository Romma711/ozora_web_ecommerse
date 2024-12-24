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

func (s *store) CreateCartItem(productId int64, quantity int, price float64) (int64, error) {
	res, err := s.db.Exec(`INSERT INTO cart_shopping (product_id, quantity, price) VALUES ($1, $2, $3)`, productId, quantity, price)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (s *store) CreateCart(userId int64, total float64, address string) (int64, error) {
	res, err := s.db.Exec(`INSERT INTO cart (id_client, total, address) VALUES ($1, $2, $3, $4)`, userId, total, address)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (s *store) CreateOrder (idCart int64, idCartItem int64) error {
	_, err := s.db.Exec(`INSERT INTO order (id_cart, id_cart_item) VALUES ($1, $2)`, idCart, idCartItem)
	return err
}

