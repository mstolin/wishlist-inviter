package db

import (
	"fmt"

	"github.com/mstolin/present-roulette/utils/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var ErrNoMatch = fmt.Errorf("no matching record")

type DatabaseHandler struct {
	DB *gorm.DB
}

func Initialize(host, port, user, password, dbName string) (DatabaseHandler, error) {
	handler := DatabaseHandler{}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return handler, err
	}

	// auto migrate models
	// TODO put in own method
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Item{})

	handler.DB = db
	return handler, nil
}
