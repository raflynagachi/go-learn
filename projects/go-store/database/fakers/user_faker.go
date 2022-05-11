package fakers

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/raflynagachi/go-store/models"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	return &models.User{
		ID:            uuid.NewString(),
		Addresses:     []models.Address{},
		Firstname:     faker.FirstName(),
		Lastname:      faker.LastName(),
		Email:         faker.Email(),
		Password:      "password",
		RememberToken: "",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     gorm.DeletedAt{},
	}
}
