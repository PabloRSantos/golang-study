package config

import (
	"fmt"
	model "go-api/app/domain/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDb() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Conected to " + dbname)

	db.AutoMigrate(&model.User{}, &model.Event{}, &model.Subscription{})
	db.SetupJoinTable(&model.User{}, "Events", &model.Subscription{})
	db.SetupJoinTable(&model.Event{}, "Users", &model.Subscription{})

	return db
}
