package models

type Ingredient struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255"`
	Category string `gorm:"size:255"`
}
