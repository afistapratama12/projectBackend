package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Config() *gorm.DB {
	var dbUser, dbPass, dbHost, dbName string

	if dbUser = os.Getenv("DB_USER"); dbUser == "" {
		dbUser = "root"
	}

	if dbPass = os.Getenv("DB_PASS"); dbPass == "" {
		dbPass = ""
	}

	if dbHost = os.Getenv("DB_HOST"); dbHost == "" {
		dbHost = "localhost"
	}

	if dbName = os.Getenv("DB_NAME"); dbName == "" {
		dbName = "refactory_akselerasi"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// db.AutoMigrate(&user.User{})
	// db.AutoMigrate(&note.Note{})

	return db
}
