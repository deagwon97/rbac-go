package checker

import (
	"rbac-go/database"
	"rbac-go/rbac/dblayer"
	"rbac-go/rbac/models"
	// "github.com/gin-gonic/gin"
)

type RBAC struct {
	Permissions []models.Permission
	db          dblayer.DBLayer
}

type RBACInterface interface {
	CheckPermission(
		subjectID int,
		permissionServiceName string,
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
	var permissions []models.Permission
	rbac := &RBAC{
		Permissions: permissions,
		db:          db,
	}
	return rbac
}

func (rbac *RBAC) CheckPermission(
	subjectID int,
	permissionServiceName string,
	permissionName string,
	permissionAction string,
) (
	isAllowed bool,
	objects []string,
	err error,
) {
	objects, err = rbac.db.GetAllowedObjects(
		subjectID,
		permissionServiceName,
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
	ServiceName string,
	Name string,
	Actions []string,
	Objects []string,
) {

	for _, action := range Actions {
		for _, object := range Objects {
			permission := models.Permission{
				ServiceName: ServiceName,
				Name:        Name,
				Action:      action,
				Object:      object,
			}
			rbac.Permissions = append(
				rbac.Permissions,
				permission)
		}
	}

}

var Rbac = NewRBAC()
