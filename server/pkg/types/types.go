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
	GetProducts() ([]Product, error)
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
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Tipos
type Type struct {
	ID   int    `json:"id"`
	Name string `json:"name" `
}

//Obras
type ArtWork struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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
	ID        int       `json:"id"`
	CreatedAt string    `json:"created_at"`
	UserID    int       `json:"user_id" `
	Total     float64   `json:"total" `
	Product   []Product `json:"product"`
}

type CartItem struct {
	ID        int    `json:"id"`
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
	Token     string    `json:"token"`
	Address   string    `json:"address"`
	Productid []int     `json:"product_id"`
	Quantity  []int     `json:"quantity"`
	Price     []float64 `json:"price"`
}

type CartStore interface {
	CreateCartItem(productId int64, quantity int, price float64) (int64, error)
	CreateCart(userId int64, total float64, address string) (int64, error)
	CreateOrder(idCart int64, idCartItem int64) error
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
