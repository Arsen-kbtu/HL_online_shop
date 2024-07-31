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

	db.Table("orders_shop").AutoMigrate(&Order{})
}

func GetAllOrdersRepo() ([]Order, error) {
	var orders []Order
	result := db.Find(&orders)
	return orders, result.Error
}

func GetOrderByIDRepo(id uint) (*Order, error) {
	var order Order
	result := db.First(&order, id)
	return &order, result.Error
}

func CreateOrderRepo(order *Order) error {
	result := db.Create(order)
	return result.Error
}

func UpdateOrderRepo(order *Order) error {
	result := db.Save(order)
	return result.Error
}

func DeleteOrderRepo(id uint) error {
	result := db.Delete(&Order{}, id)
	return result.Error
}

func SearchOrdersRepo(userID uint, status string) ([]Order, error) {
	var orders []Order
	query := db.Model(&Order{})

	if userID != 0 {
		query = query.Where("user_id = ?", userID)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	result := query.Find(&orders)
	return orders, result.Error
}
