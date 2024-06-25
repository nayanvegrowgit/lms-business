package repository

import (
	"booksMan/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func CreateBorrowRecord(br *models.BorrowingRecord) (*models.BorrowingRecord, error) {
	var book models.Book
	var result *gorm.DB
	result = Db.Where("BookID = ?", br.BookID).Find(&book)
	if result.Error == nil && book.Available > 0 {
		result = Db.Create(br)
		if result.Error != nil {
			result = Db.Model(&book).Updates(models.Book{Available: book.Available - 1})
		}
	}
	return br, result.Error
}

func SearchBorrowRecord(brs []models.BorrowingRecord, userID uint) (uint64, error) {
	s := fmt.Sprintf("%d", userID)
	result := Db.Where("UserID = ?", s).Find(&brs)
	return uint64(result.RowsAffected), result.Error
}

func UpdateBorrowRecord(id uint) (models.BorrowingRecord, error) { // Retur Book
	var br models.BorrowingRecord
	result := Db.First(&br, id)
	if result.Error != nil {
		result := Db.Model(&br).Updates(models.BorrowingRecord{DateOfReturn: time.Now().Format("2002-01-01")})
		return br, result.Error
	}
	return br, nil
}

func AllBorrowRecord(brs []models.BorrowingRecord) (int64, error) {
	result := Db.Find(&brs)  // SELECT * FROM users;
	if result.Error != nil { // returns error
		return 0, result.Error
	} else {
		return result.RowsAffected, result.Error
	}
}

func DeleteBorrowRecord(id uint) error {
	var br models.BorrowingRecord
	result := Db.Delete(&br, id)
	return result.Error
}
