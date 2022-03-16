package dblayer

import "rbac-go/account/models"

type DBLayer interface {
	AddUser(models.User) (models.User, error)
	GetPassword(string) (models.User, error)
	GetUserListName([]int) ([]UserIDName, error)
}
