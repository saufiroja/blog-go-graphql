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

func InitDatabase(conf *Config) *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf.DBUSER = os.Getenv("DB_USER")
	conf.DBPASS = os.Getenv("DB_PASS")
	conf.DBNAME = os.Getenv("DB_NAME")
	conf.DBHOST = os.Getenv("DB_HOST")
	conf.DBPORT = os.Getenv("DB_PORT")
	env := os.Getenv("GO_ENV")

	var dbURI string
	if env == "development" {
		dbURI = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", conf.DBHOST, conf.DBUSER, conf.DBPASS, conf.DBNAME, conf.DBPORT)
	} else {
		dbURI = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", conf.DBHOST, conf.DBUSER, conf.DBPASS, conf.DBNAME, conf.DBPORT)
	}

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Article{}, &entity.Comment{}, &entity.Like{}, &entity.Dislike{})

	return db
}
