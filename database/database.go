package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jicodes/go-fullstack-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB // the struct has a Db filed the type is a pointer to gorm.DB
}

var DB Dbinstance // the variable DB is of type Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", 
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), 
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database!\n", err)
		os.Exit(2)
	}

	log.Printf("Connected to database %s successfully!\n", os.Getenv("DB_NAME"))
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations...")
	db.AutoMigrate(&models.Fact{}) // run migrations to create tables we need from Gorm models

	DB = Dbinstance{
		Db: db,
	} // assign the db connection to the DB variable
}