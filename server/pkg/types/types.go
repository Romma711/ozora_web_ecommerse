package types

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	BarCode     string  `json:"code" gorm:"not null, primaryKey"`
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"not null"`
	Image       string  `json:"image"`
	Category    int     `json:"category" gorm:"foreignKey:Category"`
	Type        int     `json:"type" gorm:"foreignKey:Type"`
	ArtWork     int     `json:"artwork" gorm:"foreignKey:ArtWork,"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at"`
	Sold        int     `json:"sold"`
	Stock       int     `json:"stock"`
}

type Category struct {
	gorm.Model
	Name       string `json:"name" gorm:"not null, unique"`
}

type Type struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null, unique"`
}

type ArtWork struct {
	gorm.Model
	Title     string `json:"title" gorm:"not null"`
}
