package repository

import (
	"booksMan/models"
	"fmt"
)

func POSTcreateBook(book *models.Book) (*models.Book, error) {
	result := Db.Create(book) // returns inserted data's primary key
	if result.Error != nil {
		return book, result.Error
	} else {
		return book, nil
	}
}
func POSTsearchBook(books []models.Book, searchTitle string) (uint64, error) {
	// LIKE
	s := fmt.Sprintf("%%%s%%", searchTitle)
	result := Db.Where("title LIKE ?", s).Find(&books)
	if result.Error != nil {
		return 0, result.Error
	}
	return uint64(result.RowsAffected), result.Error
}

func PATCHupdateBook(book *models.Book) error {
	Db.First(book.ID)
	result := Db.Model(&models.Book{}).Where("id = ?", book.ID).Updates(book)
	return result.Error
}

func DELETEdeleteBook(id uint) error {
	var book models.Book
	result := Db.Delete(&book, id)
	return result.Error
}

func GETreturnAllBooks() ([]models.Book, error) {
	var books []models.Book
	result := Db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	} else {
		//fmt.Printf("in db handler :: books :\n %v", books)
		return books, result.Error
	}
}
