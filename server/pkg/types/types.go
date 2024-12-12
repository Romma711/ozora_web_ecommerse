package types

///Tipos de datos
//Productos
type Product struct {
	ID          int     `json:"id"`
	BarCode     string  `json:"code"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	CategoryID  int     `json:"category"`
	TypeID      int     `json:"type"`
	ArtWorkID   int     `json:"artwork"`
	CreatedAt   string  `json:"created_at"`
	DeletedAt   string  `json:"deleted_at"`
	Status      string  `json:"status"`
	Sold        int     `json:"sold"`
	Stock       int     `json:"stock"`
}
type ProductStore interface {
	CreateProduct(product *ProductPayLoad) error
	GetProductByID(id int) (*Product, error)
	GetProductsByCategory(categoryId int) ([]Product, error)
	GetProductsByTypes(typesId int) ([]Product, error)
	GetProductsByArtWork(artWorkId int) ([]Product, error)
	UpdateProduct(product *Product) error
	DeleteProduct(id int) error
	GetProducts() ([]Product, error)
	GetProductsByStock(stock int) ([]Product, error)
}

type ProductPayLoad struct {
	BarCode     string  `json:"code"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	CategoryID  int     `json:"category"`
	TypeID      int     `json:"type"`
	ArtWorkID   int     `json:"artwork"`
	Stock       int     `json:"stock"`
}

///Tags
//Categorias
type Category struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
}
//Tipos
type Type struct {
	ID        int    `json:"id"`
	Name      string `json:"name" `
}
//Obras
type ArtWork struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
}
type Tag struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
}

type TagsStore interface {
	CreateCategory(category *Category) error
	GetCategories() ([]Category, error)
	CreateType(type_ *Type) error
	GetTypes() ([]Type, error)
	CreateArtWork(artWork *ArtWork) error
	GetArtWorks() ([]ArtWork, error)
}
///Ordenes
type Order struct {
	ID        int         `json:"id"`
	CreatedAt string      `json:"created_at"`
	DeletedAt string      `json:"deleted_at"`
	UserID    int         `json:"user_id" `
	Total     float64     `json:"total" `
	OrderItem []OrderItem `json:"order_item"`
}

type OrderItem struct {
	ID        int     `json:"id"`
	CreatedAt string  `json:"created_at"`
	DeletedAt string  `json:"deleted_at"`
	UserID    int     `json:"user_id" `
	ProductID int     `json:"product_id" `
	Quantity  int     `json:"quantity" `
	Total     float64 `json:"total" `
}

type Cart struct {
	ID        int     `json:"id"`
	CreatedAt string  `json:"created_at"`
	DeletedAt string  `json:"deleted_at"`
	UserID    int     `json:"user_id" `
	OrderID   int     `json:"product_id" `
	Total     float64 `json:"total" `
	Address   string  `json:"address" `
	Status    string  `json:"status" `
}

type User struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
	Email     string `json:"email" `
	Password  string `json:"password" `
	Role      string `json:"role" `
	Name      string `json:"name" `
	Surname   string `json:"surname" `
	Mobile    string `json:"mobile" `
}

type Login struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
