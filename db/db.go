package db

import (
	"fmt"
	"go-multirole/config"
	"go-multirole/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	// Automatically migrate schema
	db.AutoMigrate(&model.User{}, &model.Role{}, &model.Permission{})

	return db
}
