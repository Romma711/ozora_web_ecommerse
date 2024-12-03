package types

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code        string  `json:"code" gorm:"not null, primaryKey"`
	Type        string  `json:"type" gorm:"not null"`
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"not null"`
	Image       string  `json:"image"`
	Category    string  `json:"category"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at"`
	Sold        int     `json:"sold"`
	Stock       int     `json:"stock"`
}
