package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mukezhz/hexagonal-architecture/file/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDatabase(config map[string]string) (*gorm.DB, error) {
	var db *gorm.DB
	dsn := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["DB_USERNAME"],
		config["DB_PASSWORD"],
		config["DB_HOST"],
		config["DB_DATABASE"],
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.FileMetadata{}); err != nil {
		return nil, err
	}
	return db, nil
}

func NewMysqlDB() *gorm.DB {
	envConfig, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error while reading .env")
	}
	gormDB, err := InitDatabase(envConfig)
	if err != nil {
		log.Fatal("Error initializing database")
	}
	return gormDB
}
