package models

type Cuisine struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Type string `json:"type"`
	// Relations
	Menus []Menu `json:"menus"`
}
