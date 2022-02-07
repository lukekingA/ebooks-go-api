package models

type Book struct {
	// gorm.Model
	Id            int    `json:"id" gorm:"primaryKey"`
	AuthorId      int    `json:"authorId"`
	Title         string `json:"title"`
	CopyrightYear int    `json:"copyrightYear"`
	About         string `json:"about"`
}
