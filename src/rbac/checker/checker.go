package checker

import (
	"database/sql"
	"rbac-go/rbac/dblayer"
	// "github.com/gin-gonic/gin"
)

type Permission struct {
	Name   string
	Action string
	Object sql.NullString
}

func (Permission) NewPermissions(
	Name string,
	Actions []string,
	Objects []string,
) {
	//RBAC에 Permission 추가
}

type RBAC interface {
	CheckPermission(
		subjectID int,
		permissionName string,
		permissionAction string,
	) (
		isAllowed bool,
		objects []string,
		err error,
	)
}

type innerRBAC struct {
	Permissions []Permission
	db          dblayer.DBLayer
}

// 생성자
func NewRBAC() RBAC {
	var permissions []Permission
	rbac := &innerRBAC{Permissions: permissions}
	return rbac
}

func (rbac *innerRBAC) CheckPermission(
	subjectID int,
	permissionName string,
	permissionAction string,
) (
	isAllowed bool,
	objects []string,
	err error,
) {
	objects, err = rbac.db.GetObjects(subjectID, permissionName, permissionAction)
	if len(objects) > 0 {
		isAllowed = true
	} else {
		isAllowed = false
	}
	return isAllowed, objects, err
}

func (rbac *innerRBAC) AddPermission()

var Rbac = NewRBAC()
