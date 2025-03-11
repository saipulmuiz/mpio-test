package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/saipulmuiz/mpio-test/models"
	"github.com/saipulmuiz/mpio-test/pkg/serror"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (cfg *Config) InitPostgres() serror.SError {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname =%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatalf("failed connect to database %+v", err)
		return serror.NewFromError(err)
	}

	err = db.Debug().AutoMigrate(
		models.User{},
		models.Transaction{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database %+v", err)
		return serror.NewFromError(err)
	}

	if db.Migrator().HasTable(&models.User{}) {
		if err := db.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			users := []models.User{
				{Name: "User 1", Balance: 1000000},
				{Name: "User 2", Balance: 1500000},
			}
			if err := db.Create(&users).Error; err != nil {
				log.Printf("Error seeding users: %s", err)
			} else {
				log.Println("Users seeded successfully")
			}
		}
	}

	cfg.DB = db

	return nil
}
