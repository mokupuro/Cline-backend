package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func main() {
	dbConnect()
}

func dbConnect() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	protocol := "tcp(" + os.Getenv("DOMAIN") + ":" + os.Getenv("PORT") + ")"
	dbName := os.Getenv("DB_NAME") + "?charset=utf8&parseTime=true&loc=Local"
	CONNECT := dbUser + ":" + dbPass + "@" + protocol + "/" + dbName

	db, err := gorm.Open("mysql", CONNECT)
	DB = db

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("db connected: ", &DB)
}
