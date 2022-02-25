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
	GetRoleList() (
		roleList RoleList,
		err error,
	)
	GetRolesPage(
		page int,
		pageSize int,
		hostUrl string,
	) (
		rolePage RolePage,
		err error,
	)
	AddRole(
		roleData RoleData,
	) (
		role models.Role, err error,
	)
	UpdateRole(
		id int,
		rolData RoleData,
	) (
		role models.Role,
		err error,
	)
	DeleteRole(
		id int,
	) (
		role models.Role,
		err error,
	)
}
