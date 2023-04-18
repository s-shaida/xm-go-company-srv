package models

type Company struct {
	Id                int64  `json:"id" gorm:"primaryKey"`
	Name              string `json:"name" gorm:"size:255;not null;unique"`
	Description       string `json:"description" gorm:"size:3000"`
	AmountOfEmployees int64  `json:"amount" gorm:"not null"`
	Registered        string `json:"registered" gorm:"not null"`
	Type              string `json:"type" gorm:"not null"`
}
