package models

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	Address string
	Bedroom int
	Bathroom int
	Price float32
	Sqm int
	Mode string
	HomeType string
	CountryID uint
	Country Country `gorm:"foreignKey:CountryID"`
}
