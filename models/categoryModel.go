package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	id   int `gorm:"primaryKey"`
	name string
}
