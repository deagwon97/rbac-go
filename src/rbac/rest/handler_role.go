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

// @Summary Permission의 유효성 검증
// @Tags RBAC role
// @Accept json
// @Produce json
// @Param data body dblayer.PermissionOfRole true "Data"
// @Success 200 {object} dblayer.PermissionStatusOfRole
// @Router /rbac/role/permission [post]
func (h *Handler) CheckPermissionIsAllowed(c *gin.Context) {

	var permissionOfRole dblayer.PermissionOfRole

	err := c.ShouldBindJSON(&permissionOfRole)

	permissionStatusOfRole, err := h.db.CheckPermissionIsAllowed(permissionOfRole)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, permissionStatusOfRole)
}

// @Summary Subject의 유효성 검증
// @Tags RBAC role
// @Accept json
// @Produce json
// @Param data body dblayer.SubjectsOfRole true "Data"
// @Success 200 {object} dblayer.SubjectStatusOfRole
// @Router /rbac/role/subject [post]
func (h *Handler) CheckSubjectIsAllowed(c *gin.Context) {

	var subjectsOfRole dblayer.SubjectsOfRole

	err := c.ShouldBindJSON(&subjectsOfRole)

	subjectStatusOfRole, err := h.db.CheckSubjectIsAllowed(subjectsOfRole)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, subjectStatusOfRole)
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
