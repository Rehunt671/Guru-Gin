package models

import "time"

type Recipe struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	UserID      uint       `json:"user_id"`
	MenuID      uint       `json:"menu_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ImageName   string     `json:"image_name"`
	Recipe      string     `json:"recipe"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	// Relations
	User        User         `gorm:"foreignKey:UserID" json:"user"`
	Menu        Menu         `gorm:"foreignKey:MenuID" json:"-"`
	Ingredients []Ingredient `gorm:"many2many:ingredients_on_recipes;" json:"-"`
	Comments    []Comment    `json:"-"`
	Favorites   []Favorite   `json:"-"`
}
