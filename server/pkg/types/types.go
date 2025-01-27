package types

import "time"

// /Tipos de datos
// Productos
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

// /Esta es la interfaz de las funciones de product-service
type ProductStore interface {
	CreateProduct(product *ProductPayLoad) error
	GetProductByID(id int) (*Product, error)
	GetProductsByCategory(category string) ([]Product, error)
	GetProductsByTypes(typeName string) ([]Product, error)
	GetProductsByArtWork(artWork string) ([]Product, error)
	UpdateProduct(product *Product) error
	GetProducts() ([]Product, error)
	GetProductsByDatetime()([]Product, error)
	GetProductsBySold()([]Product, error)
}

// /Esta es la estructura para crear productos
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

type ProductResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Category    string  `json:"category"`
	Type        string  `json:"type"`
	ArtWork     string  `json:"artwork"`
	Stock       int     `json:"stock"`
}

// /Tags
// Categorias
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Tipos
type Type struct {
	ID   int    `json:"id"`
	Name string `json:"name" `
}

// Obras
type ArtWork struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
}
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TagsStore interface {
	CreateCategory(category *Category) error
	GetCategories() ([]Category, error)
	GetCategoryById(id int) (string, error)
	UpdateCategory(category Category) error
	CreateType(type_ *Type) error
	GetTypes() ([]Type, error)
	GetTypeById(id int) (string, error)
	UpdateType(type_ Type) error
	CreateArtWork(artWork *ArtWork) error
	GetArtWorks() ([]ArtWork, error)
	GetArtWorkById(id int) (string, error)
	GetArtWorkRecomendation(number int) (*ArtWork, error)
	UpdateArtWork(artWork ArtWork) error
}

// /Ordenes
type Order struct {
	IDCart    int     `json:"cart_id" `
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price" `
	IDProduct int     `json:"product"`
}

type OrderResponse struct {
	IDCart   int       `json:"cart_id" `
	Total    float64   `json:"total" `
	Products []Product `json:"product"`
	Quantity []int     `json:"quantity"`
}

type OrderStore interface {
	GetOrdersUndone() ([]Cart, error)
	GetOrderByOrderId(cartId int) ([]Order, error)
	GetOrdersByUserId(userId int) ([]Order, error)
}
type CartItem struct {
	ID        int    `json:"id"`
	IDCart    int    `json:"id_cart" `
	CreatedAt string `json:"created_at"`
	ProductID int    `json:"product_id" `
	Quantity  int    `json:"quantity" `
}

type Cart struct {
	ID      int     `json:"id"`
	UserID  int     `json:"user_id" `
	Total   float64 `json:"total" `
	Address string  `json:"address" `
	Status  string  `json:"status" `
}

type CartPayload struct {
	Token        string        `json:"token"`
	Address      string        `json:"address"`
	ProductsCart []ProductCart `json:"products"`
}

type ProductCart struct {
	ID       int     `json:"id"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}

type CartStore interface {
	CreateCartItem(productId int, quantity int, price float64, cartId int) error
	CreateCart(userId int, total float64, address string) (int, error)
}

// /USUARIOS
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

type UserPayLoad struct {
	Email    string `json:"email" `
	Password string `json:"password" `
	Role     string `json:"role" `
	Name     string `json:"name" `
	Surname  string `json:"surname" `
	Mobile   string `json:"mobile" `
}

type UserStore interface {
	CreateUser(user *UserPayLoad) error
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	UpdateUser(user *User) error
	GetUsers() ([]User, error)
	DeleteUser(id int) error
}

type Login struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TokenContent struct {
	ID   int       `json:"id"`
	Role string    `json:"role"`
	Name string    `json:"name"`
	Exp  time.Time `json:"exp"`
}
