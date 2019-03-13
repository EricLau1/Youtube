package models

import (
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DB_USER = "root"
	DB_PASS = "@root"
	DB_HOST = "127.0.0.1"
	DB_PORT = 3306
	DB_NAME = "mydb"
)

func Connect() *gorm.DB {
	URL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}