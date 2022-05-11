package seeders

import (
	"github.com/raflynagachi/go-store/database/fakers"
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeeders(db *gorm.DB) []Seeder {
	user := fakers.UserFaker(db)
	product := fakers.ProductFaker(db, user.ID)
	return []Seeder{
		{Seeder: user},
		{Seeder: product},
	}
}

func DBSeed(db *gorm.DB) error {
	for _, seeder := range RegisterSeeders(db) {
		err := db.Debug().Create(seeder.Seeder).Error
		if err != nil {
			return err
		}
	}
	return nil
}
