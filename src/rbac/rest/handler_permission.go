package rest

import (
	"net/http"
	ce "rbac-go/common/error"
	"rbac-go/common/paginate"
	"rbac-go/rbac/dblayer"

	"github.com/gin-gonic/gin"

	"strconv"
)

// @Summary Permission 목록 조회
// @Tags RBAC permission
// @Accept json
// @Produce json
// @Param page query int false  "Page Number"
// @Param pageSize query int false  "Page Size"
// @Success 200 {object} dblayer.PermissionsPage
// @Router /rbac/permission/list [get]
func (h *Handler) GetPermissionsPage(c *gin.Context) {
	page, pageSize, hostUrl := paginate.ParsePageUrl(c)
	permissions, err := h.db.GetPermissionsPage(page, pageSize, hostUrl)
	ce.GinError(c, err)
	c.JSON(http.StatusOK, permissions)
}

// @Summary Permission 생성
// @Tags RBAC permission
// @Accept json
// @Produce json
// @Param data body dblayer.PermissionData true "Data"
// @Success 200 {object} models.Permission
// @Router /rbac/permission [post]
func (h *Handler) AddPermission(c *gin.Context) {
	var permissionData dblayer.PermissionData

	err := c.ShouldBindJSON(&permissionData)
	ce.GinError(c, err)

	permission, err := h.db.AddPermission(permissionData)
	ce.GinError(c, err)
	c.JSON(http.StatusOK, permission)
}

// @Summary Permission Update
// @Tags RBAC permission
// @Accept json
// @Produce json
// @Param id path int  true  "Permission ID"
// @Param data body dblayer.PermissionData true "Update에 사용할 Data"
// @Success 200 {object} models.Permission "수정된 Permission 데이터"
// @Router /rbac/permission/{id} [patch]
func (h *Handler) UpdatePermission(c *gin.Context) {

	p := c.Param("id")
	id, err := strconv.Atoi(p)
	ce.GinError(c, err)
	var permissionData dblayer.PermissionData

	err = c.ShouldBindJSON(&permissionData)
	ce.GinError(c, err)

	permission, err := h.db.UpdatePermission(id, permissionData)
	ce.GinError(c, err)
	c.JSON(http.StatusOK, permission)
}

// @Summary Permission 삭제
// @Tags RBAC permission
// @Accept json
// @Produce json
// @Param id path int true  "Permission ID"
// @Success 200 {object} models.Permission "삭제된 Permission 데이터"
// @Router /rbac/permission/{id} [delete]
func (h *Handler) DeletePermission(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	ce.GinError(c, err)

	permission, err := h.db.DeletePermission(id)
	ce.GinError(c, err)
	c.JSON(http.StatusOK, permission)
}
