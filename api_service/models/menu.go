package models

type Menu struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	CuisineID uint   `json:"cuisine_id"`
	Name      string `gorm:"unique" json:"name"`
	Rating    int    `json:"rating"`
	// Relations
	Cuisine Cuisine  `gorm:"foreignKey:CuisineID" json:"cuisine"`
	Recipes []Recipe `json:"recipes"`
}
