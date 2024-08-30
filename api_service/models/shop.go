package models

type Shop struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	PhoneNumber string  `json:"phone_number"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	BaseURL     string  `json:"base_url"`
}
