package user

import "time"

type User struct {
	ID           int `json:"user_id" gorm:"column:user_id"`
	Name              string `json:"name" gorm:"column:name"`
	Occupasion        string `json:"occupasion"`
	Email             string `json:"email"`
	Password_hash     string `json:"password_hash"`
	AvatarFieldName string `json:"avatar_field_name"`
	Role              string `json:"role"`
	Token              string `json:"token"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}