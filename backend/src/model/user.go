package model

type User struct {
	ID       uint    `json:"id" gorm:"primaryKey"` //主キーになる
	Name     string  `json:"name"`
	Email    string  `json:"email" gorm:"unique"` //重複を許さない
	Password string  `json:"password"`
	IconUrl  *string `json:"icon_url"`
}

type UserResponse struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	Name    string  `json:"name"`
	Email   string  `json:"email" gorm:"unique"`
	IconUrl *string `json:"icon_url"`
}
