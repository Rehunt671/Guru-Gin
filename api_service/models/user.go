package models

type User struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	FirstName    string  `json:"first_name"`
	MiddleName   string  `json:"middle_name"`
	LastName     string  `json:"last_name"`
	PhoneNumber  string  `json:"phone_number"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	AddressLine1 string  `json:"address_line_1"`
	AddressLine2 string  `json:"address_line_2"`
	// Relations
	Recipes         []Recipe        `json:"recipes"`
	SearchHistories []SearchHistory `json:"search_histories"`
	Favorites       []Favorite      `json:"favorites"`
	Notifications   []Notification  `json:"notifications"`
	UserCartItems   []UserCart      `json:"user_cart_items"`
}
