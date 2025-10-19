package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/naufal225/go-blog/models"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/go_blog_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal konek database: %v", err)
	}

	db.AutoMigrate(&models.User{}, &models.Post{})

	DB = db

	fmt.Println("Berhasil konek ke database")
}