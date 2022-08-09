package user

import "time"

type User struct {
	ID           int `json:"user_id" gorm:"column:user_id"`
	Name              string `json:"name" gorm:"column:name"`
	Occupasion        string
	Email             string
	Password_hash     string
	AvatarFieldName string
	Role              string
	Token              string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}