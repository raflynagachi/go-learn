package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID               string     `gorm:"size:36;not null;uniqueIndex;primaryKey"`
	ParentID         string     `gorm:"size:36;index"`
	Categories       []Category `gorm:"many2many:product_categories"`
	User             User
	UserID           string          `gorm:"size:36;index"`
	Sku              string          `gorm:"size:100;index"`
	Name             string          `gorm:"size:255"`
	Slug             string          `gorm:"size:255"`
	Price            decimal.Decimal `gorm:"type:decimal(16,2)"`
	Stock            int
	Weight           decimal.Decimal `gorm:"type:decimal(10,2)"`
	ShortDescription string          `gorm:"size:255"`
	Description      string          `gorm:"type:text"`
	Status           int             `gorm:"default:0"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	ProductImages    []ProductImage
}
