package models

type LocationResponse struct {
	Status   int      `json:"status"`
	Location Location `json:"location"`
}

type LocationsResponse struct {
	Status     int        `json:"status"`
	Page       int        `json:"page"`
	MaxPerPage int        `json:"maxPerPage"`
	Total      int64      `json:"total"`
	Locations  []Location `json:"locations"`
}

type Location struct {
	Name   string `gorm:"type:varchar(255);primary_key" json:"name"`
	System string `gorm:"type:varchar(255)" json:"system"`
	Planet string `gorm:"type:varchar(255)" json:"planet"`
	Moon   string `gorm:"type:varchar(255)" json:"moon"`
	Place  string `gorm:"type:varchar(255)" json:"place"`
}
