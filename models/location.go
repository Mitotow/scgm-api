package models

type Location struct {
	Name   string `gorm:"type:varchar(255);primary_key" json:"name"`
	System string `gorm:"type:varchar(255)" json:"system"`
	Planet string `gorm:"type:varchar(255)" json:"planet"`
	Moon   string `gorm:"type:varchar(255)" json:"moon"`
	Place  string `gorm:"type:varchar(255)" json:"place"`
}
