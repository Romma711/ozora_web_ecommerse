package product

import (
	"database/sql"
	"fmt"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
)

type StoreDB struct {
	db *sql.DB
}

func NewStoreDB(db *sql.DB) *StoreDB {
	return &StoreDB{db: db}
}

func (DB *StoreDB) GetProducts() ([]types.Product, error) {
	var products []types.Product

	rows, err := DB.db.Query("SELECT id, barcode, name, description, price, image, category_id, type_id, artwork_id, stock FROM products")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}

	return products, nil
}

func (DB *StoreDB) GetProductByID(id int) (*types.Product, error) {

	rows, err := DB.db.Query("SELECT id, barcode, name, description, price, image, category_id, type_id, artwork_id, stock FROM products WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	product := new(types.Product)

	for rows.Next() {
		product, err = scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}
	if product.ID == 0 {
		return nil, fmt.Errorf("Product not found")
	}
	return product, nil
}

func (DB *StoreDB) GetProductsByCategory(categoryId int) ([]types.Product, error) {
	rows, err := DB.db.Query("SELECT id, barcode, name, description, price, image, category_id, type_id, artwork_id, stock FROM products WHERE category_id = ?", categoryId)
	if err != nil {
		return nil, err
	}
	products := make([]types.Product, 0)

	for rows.Next() {
		product := new(types.Product)
		err = rows.Scan(&product.ID, &product.BarCode, &product.Name, &product.Description, &product.Price, &product.Image, &product.CategoryID, &product.TypeID, &product.ArtWorkID, &product.Status, &product.CreatedAt, &product.DeletedAt, &product.Sold, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}
	return products, nil
}

func (DB *StoreDB) GetProductsByArtWork(artWorkId int) ([]types.Product, error) {
	rows, err := DB.db.Query("SELECT id, barcode, name, description, price, image, category_id, type_id, artwork_id, stock FROM products WHERE artwork_id = ?", artWorkId)
	if err != nil {
		return nil, err
	}
	products := make([]types.Product, 0)

	for rows.Next() {
		product := new(types.Product)
		err = rows.Scan(&product.ID, &product.BarCode, &product.Name, &product.Description, &product.Price, &product.Image, &product.CategoryID, &product.TypeID, &product.ArtWorkID, &product.Status, &product.CreatedAt, &product.DeletedAt, &product.Sold, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}
	return products, nil
}

func (DB *StoreDB) GetProductsByTypes(typesId int) ([]types.Product, error) {
	rows, err := DB.db.Query("SELECT id, barcode, name, description, price, image, category_id, type_id, artwork_id, stock FROM products WHERE type_id = ?", typesId)
	if err != nil {
		return nil, err
	}
	products := make([]types.Product, 0)

	for rows.Next() {
		product := new(types.Product)
		err = rows.Scan(&product.ID, &product.BarCode, &product.Name, &product.Description, &product.Price, &product.Image, &product.CategoryID, &product.TypeID, &product.ArtWorkID, &product.Status, &product.CreatedAt, &product.DeletedAt, &product.Sold, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}
	return products, nil
}

func (DB *StoreDB) CreateProduct(product *types.ProductPayLoad) error {
	_, err := DB.db.Exec("INSERT INTO products (barcode, name, description, price, image, category_id, type_id, artwork_id, stock) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		product.BarCode, product.Name, product.Description, product.Price, product.Image,
		product.CategoryID, product.TypeID, product.ArtWorkID, product.Stock)
	if err != nil {
		return err
	}
	return nil
}

func (DB *StoreDB) UpdateProduct(product *types.Product) error {
	_, err := DB.db.Exec("UPDATE products SET barcode = ?, name = ?, description = ?, price = ?, image = ?, category_id = ?, type_id = ?, artwork_id = ?, sold = ?, stock = ? WHERE id = ?", product.BarCode, product.Name, product.Description, product.Price, product.Image, product.CategoryID, product.TypeID, product.ArtWorkID, product.Sold, product.Stock, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	err := rows.Scan(
		&product.ID,
		&product.BarCode,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Image,
		&product.CategoryID,
		&product.TypeID,
		&product.ArtWorkID,
		&product.Stock)
	if err != nil {
		return nil, err
	}
	return product, nil
}
