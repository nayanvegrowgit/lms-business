package repository

import (
	"booksMan/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DBConnSetup() {
	var err error
	dsn := os.Getenv("VEGROW_DATABASE_USERNAME") + ":" + os.Getenv("VEGROW_DATABASE_PASSWORD") + "@tcp(" + os.Getenv("127.0.0.1") + ":3306)/" + os.Getenv("VEGROW_PROJECT_DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		os.Exit(1)
	}
	Db.AutoMigrate(&models.Book{}, &models.BorrowingRecord{})
}
