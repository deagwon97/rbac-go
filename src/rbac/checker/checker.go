package checker

import (
	"rbac-go/database"
	"rbac-go/rbac/dblayer"
	// "github.com/gin-gonic/gin"
)

type Permission struct {
	Name   string
	Action string
	Object string
}

type RBAC struct {
	Permissions []Permission
	db          dblayer.DBLayer
}

type RBACInterface interface {
	CheckPermission(
		subjectID int,
		permissionName string,
		permissionAction string,
	) (
		isAllowed bool,
		objects []string,
		err error,
	)
	AddPermission(
		Name string,
		Actions []string,
		Objects []string,
	)
}

// 생성자
func NewRBAC() *RBAC {
	// DBORM 초기화
	dsn := database.DataSource
	db, _ := dblayer.NewORM("mysql", dsn)
	// permission list 초기화
	var permissions []Permission
	rbac := &RBAC{
		Permissions: permissions,
		db:          db,
	}
	return rbac
}

func (rbac *RBAC) CheckPermission(
	subjectID int,
	permissionName string,
	permissionAction string,
) (
	isAllowed bool,
	objects []string,
	err error,
) {
	objects, err = rbac.db.GetObjects(
		subjectID,
		permissionName,
		permissionAction,
	)
	if len(objects) > 0 {
		isAllowed = true
	} else {
		isAllowed = false
	}
	return isAllowed, objects, err
}

func (rbac *RBAC) AddPermission(
	Name string,
	Actions []string,
	Objects []string,
) {

	for _, action := range Actions {
		for _, object := range Objects {
			permission := Permission{
				Name:   Name,
				Action: action,
				Object: object,
			}
			rbac.Permissions = append(
				rbac.Permissions,
				permission)
		}
	}

}

var Rbac = NewRBAC()
