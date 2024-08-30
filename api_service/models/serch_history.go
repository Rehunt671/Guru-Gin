package models

type SearchHistory struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	UserID     uint   `json:"user_id"`
	RecipeID   uint   `json:"recipe_id"`
	ImageNames string `json:"image_names"`
	// Relations
	User   User   `gorm:"foreignKey:UserID" json:"user"`
	Recipe Recipe `gorm:"foreignKey:RecipeID" json:"recipe"`
}
