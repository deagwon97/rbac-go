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
	router.GET("/role/:id/permission", h.GetPermissionsStatusPage)
	router.GET("/role/:id/subject", h.GetSubjectsStatusPage)

	router.GET("/permission/list", h.GetPermissionsPage)
	router.POST("/permission/set", h.AddPermissionSets)
	router.POST("/permission", h.AddPermission)

	router.PATCH("/permission/:id", h.UpdatePermission)
	router.DELETE("/permission/:id", h.DeletePermission)
	router.POST("/permission/objects", h.GetAllowedObjects)

	router.POST("/permission-assignment", h.AddPermissionAssignment)
	router.DELETE("/permission-assignment", h.DeletePermissionAssignment)

	router.POST("/subject-assignment", h.AddSubjectAssignment)
	router.DELETE("/subject-assignment", h.DeleteSubjectAssignment)
}
