package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title           string `gorm:"type:string;not null"`
	Author          string
	Edition         uint `gorm:"default:1;"`
	Publisher       string
	PublicationDate string `gorm:"type:date;"`
	Genre           string
	Available       uint
	Total           uint
}
