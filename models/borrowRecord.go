package models

import (
	"gorm.io/gorm"
)

type BorrowingRecord struct {
	gorm.Model
	BookID       uint `gorm:"foreignKey:BookID"`
	UserID       uint
	DateOfIssue  string `gorm:"type:date;"`
	DateOfReturn string `grom:"default:null"`
}
