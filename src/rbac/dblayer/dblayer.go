package dblayer

import (
	"rbac-go/rbac/models"
)

type DBLayer interface {
	PermissionDBLayer
	RoleDBLayer
	// SubjectAssignmentDBLayer
	// PermissionAssignmentDBLayer
}

type PermissionDBLayer interface {
	GetObjects(
		subjectID int,
		permissionServiceName string,
		permissionName string,
		permissionAction string,
	) (
		objects []string,
		err error,
	)
	GetPermissions() (
		permissions []models.Permission,
		err error,
	)
	GetPermissionsPage(
		page int,
		pageSize int,
		hostUrl string,
	) (
		permissionsPage PermissionsPage,
		err error,
	)
	AddPermission(
		permissionData PermissionData,
	) (
		permission models.Permission, err error,
	)
	UpdatePermission(
		id int,
		permissionData PermissionData,
	) (
		permission models.Permission,
		err error,
	)
	DeletePermission(
		id int,
	) (
		permission models.Permission,
		err error,
	)
}

type RoleDBLayer interface {
	GetRoles() (
		roles []models.Role,
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
		roleData RoleData,
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

type SubjectAssignmentDBLayer interface {
	GetRolesInSubject(
		page int,
		pageSize int,
		hostUrl string,
	) (
		rolePage RolePage,
		err error,
	)
	GetSubjectsInRole(
		page int,
		pageSize int,
		hostUrl string,
	) (
		rolePage RolePage,
		err error,
	)
	AddSubjectAssignment(
		roleData RoleData,
	) (
		role models.Role, err error,
	)
	UpdateSubjectAssignment(
		id int,
		roleData RoleData,
	) (
		role models.Role,
		err error,
	)
	DeleteSubjectAssignment(
		id int,
	) (
		role models.Role,
		err error,
	)
}

type PermissionAssignmentDBLayer interface {
	GetRolesInPermission(
		page int,
		pageSize int,
		hostUrl string,
	) (
		rolePage RolePage,
		err error,
	)
	GetPermissionsInRole(
		page int,
		pageSize int,
		hostUrl string,
	) (
		rolePage RolePage,
		err error,
	)
	AddPermissionAssignment(
		roleData RoleData,
	) (
		role models.Role, err error,
	)
	UpdatePermissionAssignment(
		id int,
		roleData RoleData,
	) (
		role models.Role,
		err error,
	)
	DeletePermissionAssignment(
		id int,
	) (
		role models.Role,
		err error,
	)
}
