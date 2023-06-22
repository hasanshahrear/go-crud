package models

import "gorm.io/gorm"

type Sales struct {
	gorm.Model
	ID         int `gorm:"primaryKey"`
	Quantity   string
	UnitPrice  string
	TotalPrice string
	ProductID  int     `gorm:"foreignKey:ProductID;references:ID"`
	Product    Product `gorm:"foreignKey:ProductID"`
}
