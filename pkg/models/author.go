package models

type Author struct {
	// gorm.Model
	AuthorID   int    `json:"id" gorm:"primary_key; unique; column:id"`
	AuthorType string `json:"authorType"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
}
