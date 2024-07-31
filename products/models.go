package main

import (
	"time"
)

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id" readonly:"true" example:"1"`
	Name        string    `json:"name" validate:"required" example:"Laptop"`
	Description string    `json:"description" example:"A high-performance laptop"`
	Price       float64   `json:"price" validate:"required,gt=0" example:"1000.50"`
	Category    string    `json:"category" validate:"required" example:"Electronics"`
	Stock       int       `json:"stock" validate:"gte=0" example:"50"`
	CreatedAt   time.Time `json:"created_at" readonly:"true" example:"2023-07-20T15:04:05Z"`
}

func (Product) TableName() string {
	return "products_shop"
}
