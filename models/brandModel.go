package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}
