package rest

import (
	"net/http"
	ce "rbac-go/common/error"
	"rbac-go/rbac/dblayer"

	"github.com/gin-gonic/gin"
)

// @Summary PermissionAssignment 생성
// @Tags RBAC permissionAssignment
// @Accept json
// @Produce json
// @Param data body dblayer.PermissionAssignmentData true "Data"
// @Success 200 {object} models.PermissionAssignment
// @Router /rbac/permission-assignment [post]
func (h *Handler) AddPermissionAssignment(c *gin.Context) {
	var permissionAssignmentData dblayer.PermissionAssignmentData

	err := c.ShouldBindJSON(&permissionAssignmentData)
	if ce.GinError(c, err) {
		return
	}

	permissionAssignment, err := h.db.AddPermissionAssignment(permissionAssignmentData)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, permissionAssignment)
}

// @Summary PermissionAssignment 삭제
// @Tags RBAC permissionAssignment
// @Accept json
// @Produce json
// @Param data body dblayer.PermissionAssignmentData true "Data"
// @Success 200 {object} models.PermissionAssignment "삭제된 PermissionAssignment 데이터"
// @Router /rbac/permission-assignment [delete]
func (h *Handler) DeletePermissionAssignment(c *gin.Context) {

	var permissionAssignmentdata dblayer.PermissionAssignmentData

	err := c.ShouldBindJSON(&permissionAssignmentdata)
	if ce.GinError(c, err) {
		return
	}

	permissionAssignment, err := h.db.DeletePermissionAssignment(permissionAssignmentdata)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, permissionAssignment)
}
