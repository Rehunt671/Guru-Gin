package models

type Favorite struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	User   User `gorm:"foreignKey:UserID"` // BelongTo Relation
	MenuID uint
	Menu   Menu `gorm:"foreignKey:MenuID"` // BelongTo Relation
}
