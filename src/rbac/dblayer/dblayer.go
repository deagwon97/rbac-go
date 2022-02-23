package dblayer

import {
	"rbac/rbac/models"
	"database/sql"
}

type Permission struct {
	Name   string
	Action string
	Object sql.Nullstring
}

type RBAC interface {
	GetObjects(
		subjectID int, 
		permissionName string, 
		permissionAction string
	) (
		objects []string, err error
	)
	AddPermissionAssignment(
		subjectID int, 
		permission Permission
	) (
		subjectID int, 
	)
	UpdatePermissionAssignment()
		subjectID int, 
		permission Permission
	) (
		subjectID int, 
	)
	DeletePermissionAssignment()
	subjectID int, 
	) (
		subjectID int, 
	)
	AddSubjectAssignment(
		subjectID int, 
	) (
		subjectID int, 
	)
	UpdateSubjectAssignment()
	subjectID int, 
	) (
		subjectID int, 
	)
	DeleteSubjectAssignment()
	subjectID int, 
	) (
		subjectID int, 
	)
}

type innerRBAC interface {
	Permissions []Permission
}

func NewRBAC() (RBAC) {
	// 생성자
	return &innerRBAC {}
}

func (rbac *innerRBAC) GetObjects(
		subjectID int,
		permissionName string,
		permissionAction string 
	)(
		objects []string, err error	
	)
) {
	
}

func (rbac *innerRBAC) UpdatePermission() {
	
}

func (rbac *innerRBAC) DeletePermission() {
	
}