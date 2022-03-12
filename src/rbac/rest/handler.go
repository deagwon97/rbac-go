package rest

import (
	"rbac-go/rbac/dblayer"

	"github.com/gin-gonic/gin"

	"rbac-go/database"

	"rbac-go/rbac/checker"
)

type Handler struct {
	db dblayer.DBLayer
}

type HandlerInterface interface {
	GetRolesPage(c *gin.Context)
	AddRole(c *gin.Context)
	UpdateRole(c *gin.Context)
	DeleteRole(c *gin.Context)
	GetSubjectsOfRolePage(c *gin.Context)
	CheckPermissionIsAllowed(c *gin.Context)
	GetAllowedObjects(c *gin.Context)

	GetPermissionsPage(c *gin.Context)
	AddPermission(c *gin.Context)
	AddPermissionSets(c *gin.Context)
	UpdatePermission(c *gin.Context)
	DeletePermission(c *gin.Context)

	AddPermissionAssignment(c *gin.Context)
	UpdatePermissionAssignment(c *gin.Context)
	DeletePermissionAssignment(c *gin.Context)

	AddSubjectAssignment(c *gin.Context)
	UpdateSubjectAssignment(c *gin.Context)
	DeleteSubjectAssignment(c *gin.Context)
}

// HandlerInterface의 생성자
func NewHandler() (h HandlerInterface, err error) {

	// RBAC 초기화
	checker.NewRBAC()
	// DBORM 초기화
	dsn := database.DataSource
	db, err := dblayer.NewORM("mysql", dsn)
	if err != nil {
		return
	}
	h = &Handler{db: db}
	return
}
