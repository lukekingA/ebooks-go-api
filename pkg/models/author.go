package models

type Author struct {
	Id         int    `json:"id" gorm:"primaryKey"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
}
