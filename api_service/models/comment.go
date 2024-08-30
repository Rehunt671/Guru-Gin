package models

type Comment struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	RecipeID uint   `json:"recipe_id"`
	Comment  string `json:"comment"`
	// Relations
	Recipe Recipe `gorm:"foreignKey:RecipeID" json:"recipe"`
}
