package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
	url := os.Getenv("DATABASE_URL")
	dsn := url
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	db.Table("payments_shop").AutoMigrate(&Payment{})
}

func GetAllPaymentsRepo() ([]Payment, error) {
	var payments []Payment
	result := db.Find(&payments)
	return payments, result.Error
}

func GetPaymentByIDRepo(id uint) (*Payment, error) {
	var payment Payment
	result := db.First(&payment, id)
	return &payment, result.Error
}

func CreatePaymentRepo(payment *Payment) error {
	result := db.Create(payment)
	return result.Error
}

func UpdatePaymentRepo(payment *Payment) error {
	result := db.Save(payment)
	return result.Error
}

func DeletePaymentRepo(id uint) error {
	result := db.Delete(&Payment{}, id)
	return result.Error
}

func SearchPaymentsRepo(userID, orderID uint, status string) ([]Payment, error) {
	var payments []Payment
	query := db.Model(&Payment{})

	if userID != 0 {
		query = query.Where("user_id = ?", userID)
	}

	if orderID != 0 {
		query = query.Where("order_id = ?", orderID)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	result := query.Find(&payments)
	return payments, result.Error
}
