package rest

import (
	"net/http"
	ce "rbac-go/common/error"
	"rbac-go/rbac/dblayer"

	"github.com/gin-gonic/gin"

	"strconv"
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

// @Summary PermissionAssignment Update
// @Tags RBAC permissionAssignment
// @Accept json
// @Produce json
// @Param id path int  true  "PermissionAssignment ID"
// @Param data body dblayer.PermissionAssignmentData true "Update에 사용할 Data"
// @Success 200 {object} models.PermissionAssignment "수정된 PermissionAssignment 데이터"
// @Router /rbac/permission-assignment/{id} [patch]
func (h *Handler) UpdatePermissionAssignment(c *gin.Context) {

	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if ce.GinError(c, err) {
		return
	}
	var permissionAssignmentData dblayer.PermissionAssignmentData

	err = c.ShouldBindJSON(&permissionAssignmentData)
	if ce.GinError(c, err) {
		return
	}

	permissionAssignment, err := h.db.UpdatePermissionAssignment(id, permissionAssignmentData)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, permissionAssignment)
}

// @Summary PermissionAssignment 삭제
// @Tags RBAC permissionAssignment
// @Accept json
// @Produce json
// @Param id path int true  "PermissionAssignment ID"
// @Success 200 {object} models.PermissionAssignment "삭제된 PermissionAssignment 데이터"
// @Router /rbac/permission-assignment/{id} [delete]
func (h *Handler) DeletePermissionAssignment(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if ce.GinError(c, err) {
		return
	}

	permissionAssignment, err := h.db.DeletePermissionAssignment(id)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, permissionAssignment)
}
