package dblayer

import "rbac/account/models"

type DBLayer interface {
	AddUser(models.User) (models.User, error)
	GetPassword(string) (models.User, error)
	// Login()
	// Logout()
	// IsLogIn()
	// UpdateUser()
	// DeleteUser()
	// FindLoginId()
	// FindPassword()
	// ChangePassword()
}
