package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID         int `gorm:"primaryKey"`
	Name       string
	CategoryID int      `gorm:"foreignKey:CategoryID;references:ID"`
	BrandID    int      `gorm:"foreignKey:BrandID;references:ID"`
	Brand      Brand    `gorm:"foreignKey:BrandID"`
	Category   Category `gorm:"foreignKey:CategoryID"`
}
