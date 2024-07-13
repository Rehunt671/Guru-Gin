package models

type IngredientOnMenu struct {
	ID           uint `gorm:"primaryKey"`
	MenuID       uint
	Menu         Menu `gorm:"foreignKey:MenuID"` // BelongTo Relation
	IngredientID uint
	Ingredient   Ingredient `gorm:"foreignKey:IngredientID"` // BelongTo Relation
	Quantity     float64
	UnitType     string `gorm:"size:255"`
}
