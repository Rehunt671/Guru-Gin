package models

type UserCart struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	UserID       uint    `json:"user_id"`
	IngredientID uint    `json:"ingredient_id"`
	Price        float64 `json:"price"`
	Quantity     float64 `json:"quantity"`
	// Relations
	User       User       `gorm:"foreignKey:UserID" json:"user"`
	Ingredient Ingredient `gorm:"foreignKey:IngredientID" json:"ingredient"`
}
