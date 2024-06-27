package models

import (
	"gorm.io/gorm"
)

type BorrowingRecord struct {
	gorm.Model
	BookID       uint   `gorm:"foreignKey:BookID" json:"book_id"`
	UserID       uint   `gorm:"foreignKey:BookID" json:"user_id"`
	DateOfIssue  string `json:"date_of_issue"`
	DateOfReturn string `grom:"default:null" json:"date_of_return"`
}
