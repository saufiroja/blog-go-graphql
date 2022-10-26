package database

import (
	"fmt"
	"graphql/blog-go-graphql/entity"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBUSER string
	DBPASS string
	DBNAME string
	DBHOST string
	DBPORT string
}

func InitDatabase(conf Config) *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	env := os.Getenv("GO_ENV")

	var dbURI string
	if env == "development" {
		dbURI = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	} else {
		dbURI = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=enable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	}

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	db.AutoMigrate(&entity.User{}, &entity.Article{})
	if err != nil {
		log.Panic(err)
	}

	return db
}
