package rest

import (
	"net/http"
	ce "rbac-go/common/error"
	"rbac-go/common/paginate"
	"rbac-go/rbac/dblayer"

	"github.com/gin-gonic/gin"

	"strconv"
)

// @Summary PermissionAssignment 목록 조회
// @Tags RBAC permissionAssignment
// @Accept json
// @Produce json
// @Param page query int false  "Page Number"
// @Param pageSize query int false  "Page Size"
// @Success 200 {object} dblayer.PermissionAssignmentsPage
// @Router /rbac/permissionAssignment/list [get]
func (h *Handler) GetPermissionAssignmentsPage(c *gin.Context) {
	page, pageSize, hostUrl := paginate.ParsePageUrl(c)
	permissionAssignments, err := h.db.GetPermissionAssignmentsPage(page, pageSize, hostUrl)
	ce.GinError(c, err)
	c.JSON(http.StatusOK, permissionAssignments)
}

// @Summary PermissionAssignment 생성
// @Tags RBAC permissionAssignment
// @Accept json
// @Produce json
// @Param data body dblayer.PermissionAssignmentData true "Data"
// @Success 200 {object} models.PermissionAssignment
// @Router /rbac/permissionAssignment [post]
func (h *Handler) AddPermissionAssignment(c *gin.Context) {
	var permissionAssignmentData dblayer.PermissionAssignmentData

	err := c.ShouldBindJSON(&permissionAssignmentData)
	ce.GinError(c, err)

	permissionAssignment, err := h.db.AddPermissionAssignment(permissionAssignmentData)
	ce.GinError(c, err)
	c.JSON(http.StatusOK, permissionAssignment)
}

// @Summary PermissionAssignment Update
// @Tags RBAC permissionAssignment
// @Accept json
// @Produce json
// @Param id path int  true  "PermissionAssignment ID"
// @Param data body dblayer.PermissionAssignmentData true "Update에 사용할 Data"
// @Success 200 {object} models.PermissionAssignment "수정된 PermissionAssignment 데이터"
// @Router /rbac/permissionAssignment/{id} [patch]
func (h *Handler) UpdatePermissionAssignment(c *gin.Context) {

	p := c.Param("id")
	id, err := strconv.Atoi(p)
	ce.GinError(c, err)
	var permissionAssignmentData dblayer.PermissionAssignmentData

	err = c.ShouldBindJSON(&permissionAssignmentData)
	ce.GinError(c, err)

	permissionAssignment, err := h.db.UpdatePermissionAssignment(id, permissionAssignmentData)
	ce.GinError(c, err)
	c.JSON(http.StatusOK, permissionAssignment)
}

// @Summary PermissionAssignment 삭제
// @Tags RBAC permissionAssignment
// @Accept json
// @Produce json
// @Param id path int true  "PermissionAssignment ID"
// @Success 200 {object} models.PermissionAssignment "삭제된 PermissionAssignment 데이터"
// @Router /rbac/permissionAssignment/{id} [delete]
func (h *Handler) DeletePermissionAssignment(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	ce.GinError(c, err)

	permissionAssignment, err := h.db.DeletePermissionAssignment(id)
	ce.GinError(c, err)
	c.JSON(http.StatusOK, permissionAssignment)
}
