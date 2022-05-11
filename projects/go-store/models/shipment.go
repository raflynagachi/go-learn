package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Shipment struct {
	ID          string `gorm:"size:36;not null;index;primaryKey"`
	User        User
	UserID      string `gorm:"size:36"`
	Order       Order
	OrderID     string `gorm:"size:36"`
	TrackNumber string `gorm:"size:255;index"`
	Status      string `gorm:"size:36;index"`
	TotalQty    int
	TotalWeight decimal.Decimal `gorm:"type:decimal(10,2)"`
	Firstname   string          `gorm:"size:100;not null"`
	Lastname    string          `gorm:"size:100;not null"`
	CityID      string          `gorm:"size:100"`
	ProvinceID  string          `gorm:"size:100"`
	Address1    string          `gorm:"size:100"`
	Address2    string          `gorm:"size:100"`
	Phone       string          `gorm:"size:50"`
	Email       string          `gorm:"size:100"`
	PostCode    string          `gorm:"size:100"`
	ShippedBy   string          `gorm:"size:36"`
	ShippedAt   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
