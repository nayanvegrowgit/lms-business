package models

import "time"

type Book struct {
	//	gorm.Model      `json:""`
	ID              uint      `gorm:"primarykey;" json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Title           string    `gorm:"type:string;not null" json:"title"`
	Author          string    `json:"author"`
	Edition         uint      `gorm:"default:1;" json:"edition"`
	Publisher       string    `json:"publisher"`
	PublicationDate string    `gorm:"type:date;" json:"publication_date"`
	Genre           string    `json:"genre"`
	Available       uint      `json:"available"`
	Total           uint      `json:"total"`
}
