package main

import "time"

type User struct {
	ID             uint      `gorm:"primaryKey" json:"id" readonly:"true" example:"1"`
	Name           string    `json:"name" validate:"required" example:"John Doe"`
	Email          string    `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Address        string    `json:"address" example:"123 Main St"`
	RegistrationAt time.Time `json:"registrationAt" readonly:"true" example:"2023-07-20T15:04:05Z"`
	Role           string    `json:"role" validate:"required,oneof=admin client" example:"client"`
}

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id" readonly:"true" example:"1"`
	Name        string    `json:"name" validate:"required" example:"Laptop"`
	Description string    `json:"description" example:"A high-performance laptop"`
	Price       float64   `json:"price" validate:"required,gt=0" example:"1000.50"`
	Category    string    `json:"category" validate:"required" example:"Electronics"`
	Stock       int       `json:"stock" validate:"gte=0" example:"50"`
	CreatedAt   time.Time `json:"created_at" readonly:"true" example:"2023-07-20T15:04:05Z"`
}

type Order struct {
	ID         uint      `gorm:"primaryKey" json:"id" readonly:"true" example:"1"`
	UserID     uint      `json:"user_id" validate:"required" example:"1"`
	Products   []uint    `gorm:"-" json:"products" validate:"required" `
	TotalPrice float64   `json:"total_price" validate:"required,gt=0" example:"100.50"`
	OrderDate  time.Time `json:"order_date" readonly:"true" example:"2023-07-20T15:04:05Z"`
	Status     string    `json:"status" validate:"required,oneof=new in_process completed" example:"new"`
}

type PaymentRequest struct {
	Amount     float64 `json:"amount" validate:"required" example:"100.00"`
	OrderID    int     `json:"order_id" validate:"required" example:"1"`
	UserID     int     `json:"user_id" validate:"required" example:"1"`
	HPAN       string  `json:"hpan" validate:"required" example:"4003032704547597"`
	ExpDate    string  `json:"expDate" validate:"required" example:"1022"`
	CVC        string  `json:"cvc" validate:"required" example:"636"`
	TerminalID string  `json:"terminalId" validate:"required" example:"67e34d63-102f-4bd1-898e-370781d0074d"`
}

type Payment struct {
	ID          int       `gorm:"primaryKey" json:"id" readonly:"true" example:"1"`
	UserID      int       `json:"user_id" validate:"required" example:"1"`
	OrderID     int       `json:"order_id" validate:"required" example:"1"`
	Amount      float64   `json:"amount" validate:"required,gt=0" example:"100"`
	PaymentDate time.Time `json:"payment_date" readonly:"true" example:"2023-07-20T15:04:05Z"`
	Status      string    `json:"status" validate:"required,oneof=successful unsuccessful" example:"successful"`
}
