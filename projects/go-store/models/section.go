package models

import "time"

type Section struct {
	ID         string `gorm:"size:36;not null;index;primaryKey"`
	Name       string `gorm:"size:100"`
	Slug       string `gorm:"size:100"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Categories []Category
}
