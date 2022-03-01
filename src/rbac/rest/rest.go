package rest

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/rbac")
	h, _ := NewHandler()

	router.GET("/role/list", h.GetRolesPage)
	router.POST("/role", h.AddRole)
	router.PATCH("/role/:id", h.UpdateRole)
	router.DELETE("/role/:id", h.DeleteRole)
	router.GET("/role/:id/subject", h.GetSubjectsOfRolePage)
	router.GET("/role/:id/permission", h.GetPermissionsOfRolePage)

	router.GET("/permission/list", h.GetPermissionsPage)
	router.POST("/permission", h.AddPermission)
	router.PATCH("/permission/:id", h.UpdatePermission)
	router.DELETE("/permission/:id", h.DeletePermission)
	router.POST("/permission/objects", h.GetAllowedObjects)

	router.POST("/permission-assignment", h.AddPermissionAssignment)
	router.PATCH("/permission-assignment/:id", h.UpdatePermissionAssignment)
	router.DELETE("/permission-assignment/:id", h.DeletePermissionAssignment)

	router.POST("/subject-assignment", h.AddSubjectAssignment)
	router.PATCH("/subject-assignment/:id", h.UpdateSubjectAssignment)
	router.DELETE("/subject-assignment/:id", h.DeleteSubjectAssignment)
}
