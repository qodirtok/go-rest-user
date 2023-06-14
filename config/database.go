package config

import (
	"fmt"
	"go-rest-user/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databasename := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, databasename, port)

	model.Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection Database Success")
		Migration()
	}
}

func Migration() {
	model.Database.AutoMigrate(&model.UserAttribute{})
}
