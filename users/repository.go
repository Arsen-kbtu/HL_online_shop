package main

import (
	_ "HL_online_shop/docs"
	"database/sql"
	"fmt"
)

// var db *gorm.DB
//
//	func InitDB() {
//		if err := godotenv.Load(); err != nil {
//			log.Println("Error loading .env file")
//		}
//		url := os.Getenv("DATABASE_URL")
//		dsn := url
//		var err error
//		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
//		if err != nil {
//			log.Fatal("failed to connect to the database:", err)
//		}
//
//		db.Table("users_shop").AutoMigrate(&User{})
//
// }
//
//	func GetAllUsersRepo() ([]User, error) {
//		var users []User
//		result := db.Find(&users)
//		return users, result.Error
//	}
//
//	func GetUserByIDRepo(id uint) (*User, error) {
//		var user User
//		result := db.First(&user, id)
//		return &user, result.Error
//	}
//
//	func CreateUserRepo(user *User) error {
//		result := db.Create(user)
//		return result.Error
//	}
//
//	func UpdateUserRepo(user *User) error {
//		result := db.Save(user)
//		return result.Error
//	}
//
//	func DeleteUserRepo(id uint) error {
//		result := db.Delete(&User{}, id)
//		return result.Error
//	}
//
//	func SearchUsersRepo(name, email string) ([]User, error) {
//		var users []User
//		query := db.Model(&User{})
//
//		if name != "" {
//			query = query.Where("name ILIKE ?", "%"+name+"%")
//		}
//
//		if email != "" {
//			query = query.Where("email ILIKE ?", "%"+email+"%")
//		}
//
//		result := query.Find(&users)
//		return users, result.Error
//	}
type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn string
	}
}

var db *sql.DB

func OpenDB(cfg Config) (*sql.DB, error) {
	// Use sql.Open() to create an empty connection pool, using the DSN from the config // struct.
	var err error
	db, err = sql.Open("postgres", cfg.Db.Dsn)
	fmt.Println(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

//create table if not exists users_shop(
//id serial primary key,
//name varchar(50) not null,
//email varchar(50) not null,
//address varchar(50) not null,
//registration_at timestamp not null,
//role varchar(50) not null
//);

func GetAllUsersRepo() ([]User, error) {
	var users []User
	rows, err := db.Query("SELECT * FROM users_shop")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Address, &user.RegistrationAt, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByIDRepo(id uint) (*User, error) {
	var user User
	err := db.QueryRow("SELECT * FROM users_shop WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Address, &user.RegistrationAt, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUserRepo(user *User) error {
	_, err := db.Exec("INSERT INTO users_shop (name, email, address, registration_at, role) VALUES ($1, $2, $3, $4, $5)", user.Name, user.Email, user.Address, user.RegistrationAt, user.Role)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserRepo(user *User) error {
	_, err := db.Exec("UPDATE users_shop SET name = $1, email = $2, address = $3, registration_at = $4, role = $5 WHERE id = $6", user.Name, user.Email, user.Address, user.RegistrationAt, user.Role, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserRepo(id uint) error {
	_, err := db.Exec("DELETE FROM users_shop WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func SearchUsersRepo(name, email string) ([]User, error) {
	var users []User
	rows, err := db.Query("SELECT * FROM users_shop WHERE name ILIKE $1 AND email ILIKE $2", "%"+name+"%", "%"+email+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Address, &user.RegistrationAt, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
