package models

import (
	"database/sql"
)

type User struct {
	ID       int            `gorm:"primaryKey;column:ID;" json:"ID"`
	LoginID  string         `gorm:"column:LoginID" json:"LoginID"`
	Password string         `gorm:"column:Password" json:"Password"`
	Name     sql.NullString `gorm:"column:Name"     json:"Name" swaggertype:"string"`
	Email    sql.NullString `gorm:"column:Email"    json:"Email" swaggertype:"string"`
}

func (User) TableName() string {
	return "User"
}

type AddUserData struct {
	LoginID  string `json:"LoginID" validate:"required"`
	Password string `json:"Password" validate:"required"`
	Email    string `json:"Email"`
	Name     string `json:"Name"`
}

type LoginRequest struct {
	LoginID  string `json:"LoginID" validate:"required"`
	Password string `json:"Password" validate:"required"`
}

type LoginResult struct {
	UserID       int    `json:"UserID"`
	AccessToken  string `json:"AccessToken"`
	RefreshToken string `json:"RefreshToken"`
}
