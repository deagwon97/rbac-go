package dblayer

import (
	"rbac-go/rbac/models"
)

type DBLayer interface {
	PermissionDBLayer
	RoleDBLayer
	SubjectDBLayer
	SubjectAssignmentDBLayer
	PermissionAssignmentDBLayer
}

type PermissionDBLayer interface {
	GetAllowedObjects(
		subjectID int,
		permissionServiceName string,
		permissionName string,
		permissionAction string,
	) (
		permissionAnswer PermissionAnswer,
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

	GetPermissionsStatusPage(
		roleID int,
		page int,
		pageSize int,
		hostUrl string,
	) (
		permissionPage PermissionsStatusPage,
		err error,
	)

	AddPermission(
		permissionData PermissionData,
	) (
		permission models.Permission, err error,
	)
	AddPermissionSets(
		permissionSetData PermissionSetData,
	) (
		permissions []TempPermission, err error,
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

type SubjectDBLayer interface {
	GetSubjectsStatusPage(
		roleID int,
		page int,
		pageSize int,
		hostUrl string,
	) (
		subjectPage SubjectsStatusPage,
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
		rolesPage RolesPage,
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
	CheckPermissionIsAllowed(
		permissionOfRole PermissionOfRole,
	) (
		permissionStatusOfRole PermissionStatusOfRole,
		err error,
	)
	CheckSubjectIsAllowed(
		subjectOfRole SubjectsOfRole,
	) (
		subjectStatusOfRole SubjectStatusOfRole,
		err error,
	)
}

type SubjectAssignmentDBLayer interface {
	AddSubjectAssignment(
		SubjectAssignmentData,
	) (
		models.SubjectAssignment, error,
	)
	UpdateSubjectAssignment(
		int,
		SubjectAssignmentData,
	) (
		models.SubjectAssignment,
		error,
	)
	DeleteSubjectAssignment(
		id int,
	) (
		models.SubjectAssignment,
		error,
	)
}

type PermissionAssignmentDBLayer interface {
	AddPermissionAssignment(
		PermissionAssignmentData,
	) (
		models.PermissionAssignment,
		error,
	)
	UpdatePermissionAssignment(
		int,
		PermissionAssignmentData,
	) (
		models.PermissionAssignment,
		error,
	)
	DeletePermissionAssignment(
		int,
	) (
		models.PermissionAssignment,
		error,
	)
}
