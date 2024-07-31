package main

import (
	"time"
)

type Order struct {
	ID         uint      `gorm:"primaryKey" json:"id" readonly:"true" example:"1"`
	UserID     uint      `json:"user_id" validate:"required" example:"1"`
	Products   []uint    `gorm:"-" json:"products" validate:"required" `
	TotalPrice float64   `json:"total_price" validate:"required,gt=0" example:"100.50"`
	OrderDate  time.Time `json:"order_date" readonly:"true" example:"2023-07-20T15:04:05Z"`
	Status     string    `json:"status" validate:"required,oneof=new in_process completed" example:"new"`
}

func (Order) TableName() string {
	return "orders_shop"
}
