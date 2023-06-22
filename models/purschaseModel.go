package models

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	Quantity  string
	Price     float32
	ProductID int     `gorm:"foreignKey:ProductID;references:ID"`
	Product   Product `gorm:"foreignKey:ProductID"`
}
