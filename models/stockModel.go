package models

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	Quantity  string
	ProductID int     `gorm:"foreignKey:ProductID;references:ID"`
	Product   Product `gorm:"foreignKey:ProductID"`
}
