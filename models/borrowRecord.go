package models

import (
	"gorm.io/gorm"
)

type BorrowingRecord struct {
	gorm.Model
	BookID       uint   `gorm:"foreignKey:BookID"`
	UserID       uint   `gorm:"foreignKey:BookID"`
	DateOfIssue  string `gorm:`
	DateOfReturn string `grom:"default:null"`
}
