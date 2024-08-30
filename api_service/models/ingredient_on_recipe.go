package models

type IngredientOnRecipe struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	RecipeID     uint    `json:"recipe_id"`
	IngredientID uint    `json:"ingredient_id"`
	Quantity     float64 `json:"quantity"`
	// Relations
	Recipe     Recipe     `gorm:"foreignKey:RecipeID" json:"recipe"`
	Ingredient Ingredient `gorm:"foreignKey:IngredientID" json:"ingredient"`
}
