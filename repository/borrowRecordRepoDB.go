package repository

import (
	"booksMan/models"
	"fmt"

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

type resultFormat struct {
	Id           uint   `json:"br_id"`
	Book_ID      uint   `json:"book_id"`
	Title        string ` json:"title"`
	Author       string `json:"author"`
	DateOfIssue  string `json:"date_of_issue"`
	DateOfReturn string ` json:"date_of_return"`
}

func SearchBorrowRecord(userID uint) ([]resultFormat, error) {
	var records []resultFormat
	resp := Db.Raw("SELECT borrowing_records.*,	books.title as title, books.author as author FROM borrowing_records INNER JOIN books ON borrowing_records.book_id = books.id WHERE borrowing_records.user_id = ?;", userID).Scan(&records)

	if resp.Error != nil {
		return nil, resp.Error
	}
	fmt.Printf("Result ::%v \n", records)
	fmt.Printf("number of records : %d\n", resp.RowsAffected)
	return records, nil
}

func UpdateBorrowRecord(id uint) error { // Retur Book
	fmt.Printf("\nDB Query br id : %d\n", id)
	var br models.BorrowingRecord
	result := Db.Raw("UPDATE borrowing_records SET date_of_return = CURDATE() WHERE id = ?;", id).Scan(&br)

	if result.Error != nil {
		fmt.Printf("Error in query : %s", result.Error)
	}

	result = Db.First(&br, id)
	fmt.Printf("br : %v", br)

	var book models.Book
	result = Db.First(&book, br.BookID)
	if result.Error != nil {
		return result.Error
	}
	book.Available = book.Available + 1
	result = Db.Save(&book)
	if result.Error != nil {
		return result.Error
	}
	fmt.Printf("Rows updated in query : %d", result.RowsAffected)
	return result.Error
}

type resultallFormat struct {
	Id           uint   `json:"br_id"`
	Book_ID      uint   `json:"book_id"`
	User_ID      uint   `json:"user_id"`
	Title        string ` json:"title"`
	Author       string `json:"author"`
	DateOfIssue  string `json:"date_of_issue"`
	DateOfReturn string ` json:"date_of_return"`
}

func AllBorrowRecord() ([]resultallFormat, error) {
	var records []resultallFormat
	resp := Db.Raw("SELECT borrowing_records.*,	books.title as title, books.author as author FROM borrowing_records INNER JOIN books ON borrowing_records.book_id = books.id;").Scan(&records)

	if resp.Error != nil {
		return nil, resp.Error
	}
	fmt.Printf("Result ::%v \n", records)
	fmt.Printf("number of records : %d\n", resp.RowsAffected)
	return records, nil
}

func DeleteBorrowRecord(id uint) error {
	var br models.BorrowingRecord
	result := Db.Delete(&br, id)
	return result.Error
}
