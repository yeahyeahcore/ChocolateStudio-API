package store

import (
	"fmt"

	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func NewStore(config *models.Config) *Store {
	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s user=%s database=%s password=%s port=%s sslmode=disable",
			config.DB.Host, config.DB.User, config.DB.DBName, config.DB.Password, config.DB.Port),
	), &gorm.Config{})
	db.AutoMigrate(&models.PhotoBook{}, &models.Book_image{}, &models.Photo_image{}, &models.Photography{})
	if err != nil {
		panic(err)
	}
	return &Store{
		DB: db,
	}
}
