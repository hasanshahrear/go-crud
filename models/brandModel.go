package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	id   int `gorm:"primaryKey"`
	name string
}
