package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vrtineu/board/apps/server/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable",
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	if err = db.AutoMigrate(&domain.Board{}); err != nil {
		log.Println(err)
	}

	return db
}
