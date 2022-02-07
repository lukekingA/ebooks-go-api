package models

type Book struct {
	// gorm.Model
	ID            int      `json:"id" gorm:"primary_key;unique;column:id"`
	Author        []Author `json:"author" gorm:"foreignKey:AuthorID"`
	Title         string   `json:"title"`
	CopyrightYear int      `json:"copyrightYear"`
	About         string   `json:"about"`
}
