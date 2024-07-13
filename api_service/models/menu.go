package models

type Menu struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      User   `gorm:"foreignKey:UserID"` // BelongTo Relation
	Name      string `gorm:"size:255"`
	ImageName string `gorm:"size:255"`
	Rating    int
}
