package models

type User struct {
	ID          uint   `gorm:"primaryKey"`
	Username    string `gorm:"size:255"`
	Password    string `gorm:"size:255"`
	FirstName   string `gorm:"size:255"`
	MiddleName  string `gorm:"size:255"`
	LastName    string `gorm:"size:255"`
	PhoneNumber string `gorm:"size:255"`
}

