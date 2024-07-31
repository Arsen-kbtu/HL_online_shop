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

	db.Table("users_shop").AutoMigrate(&User{})

}

func GetAllUsersRepo() ([]User, error) {
	var users []User
	result := db.Find(&users)
	return users, result.Error
}

func GetUserByIDRepo(id uint) (*User, error) {
	var user User
	result := db.First(&user, id)
	return &user, result.Error
}

func CreateUserRepo(user *User) error {
	result := db.Create(user)
	return result.Error
}

func UpdateUserRepo(user *User) error {
	result := db.Save(user)
	return result.Error
}

func DeleteUserRepo(id uint) error {
	result := db.Delete(&User{}, id)
	return result.Error
}

func SearchUsersRepo(name, email string) ([]User, error) {
	var users []User
	query := db.Model(&User{})

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	if email != "" {
		query = query.Where("email ILIKE ?", "%"+email+"%")
	}

	result := query.Find(&users)
	return users, result.Error
}
