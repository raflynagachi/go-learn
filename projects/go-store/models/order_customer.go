package models

import (
	"time"
)

type OrderCustomer struct {
	ID         string `gorm:"size:36;not null;index;primaryKey"`
	User       User
	UserID     string `gorm:"size:36;index"`
	Order      Order
	OrderID    string `gorm:"size:36;index"`
	Firstname  string `gorm:"size:100;not null"`
	Lastname   string `gorm:"size:100;not null"`
	CityID     string `gorm:"size:100"`
	ProvinceID string `gorm:"size:100"`
	Address1   string `gorm:"size:100"`
	Address2   string `gorm:"size:100"`
	Phone      string `gorm:"size:50"`
	Email      string `gorm:"size:100"`
	PostCode   string `gorm:"size:100"`
	CreatedAt  time.Time
	UpdateAt   time.Time
}
