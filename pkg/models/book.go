package models

type Book struct {
	Id            int    `json:"id" gorm:"primaryKey"`
	Title         string `json:"title"`
	Description   string `json:"desc"`
	CopyrightYear int    `json:"copyrightYear"`
	Author        string `json:"author"`
}
