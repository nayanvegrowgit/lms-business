package repository

import (
	"booksMan/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func CreateBorrowRecord(br *models.BorrowingRecord) (*models.BorrowingRecord, error) {
	var result *gorm.DB
	var book models.Book
	result = Db.First(&book, br.BookID)
	if result.Error != nil {
		return nil, result.Error
	}
	if book.Available > 0 {
		result = Db.Create(&br)
		fmt.Println(result.Error)
		if result.Error != nil {
			return nil, result.Error
		}
		book.Available = book.Available - 1
		result = Db.Save(&book)
		//result = Db.Model(&book).Updates(models.Book{Available: book.Available - 1})
		if result.Error != nil {
			return nil, result.Error
		}
		return br, nil
	}
	return nil, nil
}

func SearchBorrowRecord(userID uint) ([]models.Book, error) {
	var books []models.Book
	type Result struct {
		id uint
	}
	var result []Result
	resp := Db.Raw("SELECT book_id FROM borrowing_records WHERE user_id = ?", userID).Scan(&result)
	if resp.Error != nil {
		return nil, resp.Error
	}

	fmt.Print("Result ::")
	fmt.Printf(" %d ", len(result))
	fmt.Print("\n")
	if len(result) != 0 {
		resp = Db.Find(&books, result)
		if resp.Error != nil || len(books) == 0 {
			return nil, resp.Error
		}
		return books, nil
	}
	println("No length of book id = 0")
	return nil, nil
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

func AllBorrowRecord() ([]models.BorrowingRecord, error) {
	var brs []models.BorrowingRecord
	result := Db.Find(&brs)  // SELECT * FROM users;
	if result.Error != nil { // returns error
		return nil, result.Error
	} else {
		return brs, result.Error
	}
}

func DeleteBorrowRecord(id uint) error {
	var br models.BorrowingRecord
	result := Db.Delete(&br, id)
	return result.Error
}
