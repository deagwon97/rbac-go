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

	router.GET("/permission/list", h.GetPermissionsPage)
	router.POST("/permission", h.AddPermission)
	router.PATCH("/permission/:id", h.UpdatePermission)
	router.DELETE("/permission/:id", h.DeletePermission)

	router.GET("/permissionAssignment/list", h.GetPermissionAssignmentsPage)
	router.POST("/permissionAssignment", h.AddPermissionAssignment)
	router.PATCH("/permissionAssignment/:id", h.UpdatePermissionAssignment)
	router.DELETE("/permissionAssignment/:id", h.DeletePermissionAssignment)
}
