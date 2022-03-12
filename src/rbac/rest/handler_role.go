package rest

import (
	"net/http"
	ce "rbac-go/common/error"
	"rbac-go/common/paginate"
	"rbac-go/rbac/dblayer"

	"github.com/gin-gonic/gin"

	"strconv"
)

// @Summary Role 목록 조회
// @Tags RBAC role
// @Accept json
// @Produce json
// @Param page query int false  "Page Number"
// @Param pageSize query int false  "Page Size"
// @Success 200 {object} dblayer.RolesPage
// @Router /rbac/role/list [get]
func (h *Handler) GetRolesPage(c *gin.Context) {
	page, pageSize, hostUrl := paginate.ParsePageUrl(c)
	roles, err := h.db.GetRolesPage(page, pageSize, hostUrl)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, roles)
}

// @Summary 특정 Role Permission 목록 조회
// @Tags RBAC role
// @Accept json
// @Produce json
// @Param id path int true  "Role ID"
// @Param page query int false  "Page Number"
// @Param pageSize query int false  "Page Size"
// @Success 200 {object} dblayer.PermissionsPage
// @Router /rbac/role/{id}/permission [get]
func (h *Handler) GetPermissionsStatusPage(c *gin.Context) {
	p := c.Param("id")
	roleId, err := strconv.Atoi(p)
	if ce.GinError(c, err) {
		return
	}

	page, pageSize, hostUrl := paginate.ParsePageUrl(c)
	permissions, err := h.db.GetPermissionsStatusPage(roleId, page, pageSize, hostUrl)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, permissions)
}

// @Summary 특정 Role Subject 목록 조회
// @Tags RBAC role
// @Accept json
// @Produce json
// @Param id path int true  "Role ID"
// @Param page query int false  "Page Number"
// @Param pageSize query int false  "Page Size"
// @Success 200 {object} dblayer.SubjectsStatusPage
// @Router /rbac/role/{id}/subject [get]
func (h *Handler) GetSubjectsStatusPage(c *gin.Context) {
	p := c.Param("id")
	roleId, err := strconv.Atoi(p)
	if ce.GinError(c, err) {
		return
	}

	page, pageSize, hostUrl := paginate.ParsePageUrl(c)
	subjects, err := h.db.GetSubjectsStatusPage(roleId, page, pageSize, hostUrl)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, subjects)
}

// @Summary Role 생성
// @Tags RBAC role
// @Accept json
// @Produce json
// @Param data body dblayer.RoleData true "Data"
// @Success 200 {object} models.Role
// @Router /rbac/role [post]
func (h *Handler) AddRole(c *gin.Context) {
	var roleData dblayer.RoleData

	err := c.ShouldBindJSON(&roleData)
	if ce.GinError(c, err) {
		return
	}

	role, err := h.db.AddRole(roleData)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, role)
}

// @Summary Role Update
// @Tags RBAC role
// @Accept json
// @Produce json
// @Param id path int  true  "Role ID"
// @Param data body dblayer.RoleData true "Update에 사용할 Data"
// @Success 200 {object} models.Role "수정된 Role 데이터"
// @Router /rbac/role/{id} [patch]
func (h *Handler) UpdateRole(c *gin.Context) {

	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if ce.GinError(c, err) {
		return
	}
	var roleData dblayer.RoleData

	err = c.ShouldBindJSON(&roleData)
	if ce.GinError(c, err) {
		return
	}

	role, err := h.db.UpdateRole(id, roleData)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, role)
}

// @Summary Role 삭제
// @Tags RBAC role
// @Accept json
// @Produce json
// @Param id path int true  "Role ID"
// @Success 200 {object} models.Role "삭제된 Role 데이터"
// @Router /rbac/role/{id} [delete]
func (h *Handler) DeleteRole(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if ce.GinError(c, err) {
		return
	}

	role, err := h.db.DeleteRole(id)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, role)
}
