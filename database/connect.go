package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/SaiHLu/fiber-jwt/common/config"
	"github.com/SaiHLu/fiber-jwt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	p := config.Config("POSTGRES_PORT")
	port, err := strconv.Atoi(p)
	if err != nil {
		log.Fatal("DB port errors: ", err.Error())
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("POSTGRES_HOST"),
		port,
		config.Config("POSTGRES_USER"),
		config.Config("POSTGRES_PASSWORD"),
		config.Config("POSTGRES_DB"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Can't connect to database: ", err.Error())
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&models.User{})
	fmt.Println("Database Migrated")

}
