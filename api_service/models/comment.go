package models

type Comment struct {
	ID      uint `gorm:"primaryKey"`
	MenuID  uint
	Menu    Menu   `gorm:"foreignKey:MenuID"` //BelongTo Relation
	Comment string `gorm:"size:255"`
}
