package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB

	err     error
	dirver  string
	connect string
)

func init() {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Main func --> Error loading .env file")
	}
	dirver = os.Getenv("DRIVER")
	connect = os.Getenv("CONNECT")
}

// Conenect -- db Conection
func Conenect() *gorm.DB {
	db, err = gorm.Open(dirver, connect)
	if err != nil {
		panic(err.Error())
	}

	return db
}
