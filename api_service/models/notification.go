package models

import "time"

type Notification struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    UserID    uint      `json:"user_id"`
    Message   string    `json:"message"`
    IsRead    bool      `gorm:"default:false" json:"is_read"`
    CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    // Relations
    User User `gorm:"foreignKey:UserID" json:"user"`
}
