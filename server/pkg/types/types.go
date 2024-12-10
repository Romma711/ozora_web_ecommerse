package types

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
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at"`
	DeletedAt   string  `json:"deleted_at"`
	Sold        int     `json:"sold"`
	Stock       int     `json:"stock"`
}

type Category struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
	Name      string `json:"name"`
}

type Type struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
	Name      string `json:"name" `
}

type ArtWork struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
	Title     string `json:"title"`
}

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
