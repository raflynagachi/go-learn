package fakers

import (
	"math"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/raflynagachi/go-store/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func ProductFaker(db *gorm.DB, userID string) *models.Product {
	name := faker.Name() + "stuff"
	return &models.Product{
		ID:               uuid.NewString(),
		UserID:           userID,
		Sku:              slug.Make(name),
		Name:             name,
		Slug:             slug.Make(name),
		Price:            decimal.NewFromFloat(fakePrice()),
		Stock:            rand.Intn(100),
		Weight:           decimal.NewFromFloat(rand.Float64()),
		ShortDescription: name + " short description",
		Description:      faker.Paragraph(),
		Status:           1,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		DeletedAt:        gorm.DeletedAt{},
	}
}

func fakePrice() float64 {
	return precision(rand.Float64()*math.Pow10(rand.Intn(8)), rand.Intn(2)+1)
}

func precision(val float64, pre int) float64 {
	div := math.Pow10(pre)
	return float64(int64(val*div)) / div
}
