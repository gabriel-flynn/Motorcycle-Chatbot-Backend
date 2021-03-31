package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	connUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)
	conn, err := gorm.Open(mysql.Open(connUrl), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	db = conn

	db.AutoMigrate(&Motorcycle{}, &User{}, &Location{}, &Track{})
}

func GetDB() *gorm.DB {
	return db
}
