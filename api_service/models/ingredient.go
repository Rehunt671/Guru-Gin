package models

type Ingredient struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    Name      string `json:"name"`
    Category  string `json:"category"`
    UnitType  string `json:"unit_type"`
    // Relations
    Recipes []IngredientOnRecipe `json:"recipes"`
    UserCartItems []UserCart `json:"user_cart_items"`
}
