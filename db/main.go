package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	if err := ConnectDB(); err != nil {
		panic(err)
	}

	addUsers()
}

type User struct {
	ID   uint32
	Name string
}

func addUsers() {
	var users []User
	for i := range 1_0000_0000 {
		users = append(users, User{
			Name: fmt.Sprintf("user%d", i),
		})
	}

	result := db.CreateInBatches(&users, 1000)
	if result.Error != nil {
		panic(result.Error)
	}
}

var db *gorm.DB

func ConnectDB() error {
	user := "scrapt"
	passwd := "scrapt"
	host := "192.168.0.204"
	port := 3306
	dbname := "scrapt"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user, passwd, host, port, dbname)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// db.Logger = db.Logger.LogMode(logger.Info)
	db.Logger = db.Logger.LogMode(logger.Error)

	return nil
}
