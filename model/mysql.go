package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Err error

func InitDB() {
	host := "127.0.0.1"
	port := "3306"
	database := "douyin"
	username := "test1"
	password := "A123456,"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)
	Db, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if Err != nil {
		panic("Error to DB connection, err: " + Err.Error())
	}
}

func InitTable() {
	Db.AutoMigrate(&CommentData{}, &UserData{}, &VideoData{}, &FavoriteData{})
}
