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
	fmt.Printf("Update book db : book :: %v\n", book)
	result := Db.First(book.ID)
	fmt.Printf("after Db.First : result :: %v\n", result)
	fmt.Printf("book db : book :: %v\n", book)

	result = Db.Model(&models.Book{}).Where("id = ?", book.ID).Updates(book)
	fmt.Printf("after Modal : result :: %v\n", result)
	fmt.Printf("after Modal : book :: %v\n", book)

	return result.Error
}

func DELETEdeleteBook(id uint) error {
	var book models.Book
	result := Db.Delete(&book, id)
	return result.Error
}

func GETreturnAllBooks(Offset uint, Limit uint) ([]models.Book, error) {
	var books []models.Book
	rows, err := Db.Raw(" SELECT *  FROM books LIMIT ? OFFSET ?;", Limit, Offset).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = Db.ScanRows(rows, &books)
		if err != nil {
			break
		}
	}

	//	result := Db.Find(&books)
	if err != nil {
		return nil, err
	} else {
		fmt.Printf("in db handler :: books :\n %v ", books)
		return books, err
	}
}

type ValuesSum struct {
	Total     uint `json:"total"`
	Available uint `json:"available"`
}

func GetTotalAvailableBooks() ([]ValuesSum, error) {
	var value []ValuesSum
	result := Db.Raw("select sum(total) as total, sum(available) available from books;").Scan(&value)
	//result := Db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	} else {
		fmt.Printf("in db handler :: values :\n %v", value)
		return value, nil
	}
}
