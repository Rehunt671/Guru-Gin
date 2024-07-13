package models

type Recipe struct {
	ID     uint `gorm:"primaryKey"`
	MenuID uint
	Menu   Menu   `gorm:"foreignKey:MenuID"` // BelongTo Relation
	Recipe string `gorm:"size:255"`
}
