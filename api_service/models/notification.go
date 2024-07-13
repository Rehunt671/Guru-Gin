package models

type Notification struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      User   `gorm:"foreignKey:UserID"` // BelongTo Relation
	Message   string `gorm:"type:text"`
	IsRead    bool   `gorm:"default:false"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}
