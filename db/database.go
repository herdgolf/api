package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	database *gorm.DB
	e        error
)

func Init() {
	//stuff
	//
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		dbhost  = os.Getenv("DB_HOST")
		dbuser  = os.Getenv("DB_USER")
		dbpass  = os.Getenv("DB_PASSWORD")
		dbname  = os.Getenv("DB_NAME")
		dbport  = os.Getenv("DB_PORT")
		sslmode = os.Getenv("DB_SSLMODE")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbhost, dbport, dbuser, dbpass, dbname, sslmode)

	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if e != nil {
		panic(e)
	}
}

func DB() *gorm.DB {
	return database
}
