package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	db.Table("products_shop").AutoMigrate(&Product{})
}

func GetAllProductsRepo() ([]Product, error) {
	var products []Product
	result := db.Find(&products)
	return products, result.Error
}

func GetProductByIDRepo(id uint) (*Product, error) {
	var product Product
	result := db.First(&product, id)
	return &product, result.Error
}

func CreateProductRepo(product *Product) error {
	result := db.Create(product)
	return result.Error
}

func UpdateProductRepo(product *Product) error {
	result := db.Save(product)
	return result.Error
}

func DeleteProductRepo(id uint) error {
	result := db.Delete(&Product{}, id)
	return result.Error
}

func SearchProductsRepo(name, category string) ([]Product, error) {
	var products []Product
	query := db.Model(&Product{})

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	if category != "" {
		query = query.Where("category ILIKE ?", "%"+category+"%")
	}

	result := query.Find(&products)
	return products, result.Error
}
