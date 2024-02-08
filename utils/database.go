package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func ConnectSql() (db *Database, err error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, db_name)
	tl, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DATABSE CONNECTION : error")
		return db, err
	}

	fmt.Println("DATABSE CONNECTION : successfully")
	return &Database{tl}, nil

	// req := User{
	// 	Name:     "Admmi",
	// 	Email:    "admin@gmail.com",
	// 	Password: "admin123",
	// }
	// var login request.Login
	// login.Email = req.Email
	// _, _, exist := CheckUser(login)
	// if !exist {
	// 	_, _ = CreateUser(req)
	// }

}
