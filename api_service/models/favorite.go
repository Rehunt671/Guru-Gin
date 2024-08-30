package models

type Favorite struct {
    ID       uint   `gorm:"primaryKey" json:"id"`
    UserID   uint   `json:"user_id"`
    RecipeID uint   `json:"recipe_id"`
    // Relations
    User   User   `gorm:"foreignKey:UserID" json:"user"`
    Recipe Recipe `gorm:"foreignKey:RecipeID" json:"recipe"`
}
