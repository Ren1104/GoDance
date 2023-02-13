package controller

import (
	//"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDB() {
	// host := "localhost"
	// port := "3306"
	// database := "douyin"
	// username := "test1"
	// password := "12345678"
	// charset := "utf8"
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)
	dsn := "test1:12345678.@tcp(127.0.0.1)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error to DB connection, err: " + err.Error())
	}
}
