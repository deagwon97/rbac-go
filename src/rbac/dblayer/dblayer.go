package dblayer

import "rbac-go/rbac/models"

type DBLayer interface {
	GetObjects(
		subjectID int,
		permissionName string,
		permissionAction string,
	) (
		objects []string,
		err error,
	)
	GetRoles(int, int, string) ([]models.Role, error)
}
