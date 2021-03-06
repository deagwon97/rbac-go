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
	if ce.GinError(c, err) {
		return
	}
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
	if ce.GinError(c, err) {
		return
	}

	permission, err := h.db.AddPermission(permissionData)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, permission)
}

// @Summary Permission Set 생성 및 동기화
// @Tags RBAC permission
// @Accept json
// @Produce json
// @Param data body dblayer.PermissionSetData true "Data"
// @Success 200 {object} []models.Permission
// @Router /rbac/permission/set [post]
func (h *Handler) AddPermissionSets(c *gin.Context) {
	var permissionSetData dblayer.PermissionSetData

	err := c.ShouldBindJSON(&permissionSetData)
	if ce.GinError(c, err) {
		return
	}

	permissions, err := h.db.AddPermissionSets(permissionSetData)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, permissions)
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

	p := c.Param("ID")
	id, err := strconv.Atoi(p)
	if ce.GinError(c, err) {
		return
	}
	var permissionData dblayer.PermissionData

	err = c.ShouldBindJSON(&permissionData)
	if ce.GinError(c, err) {
		return
	}

	permission, err := h.db.UpdatePermission(id, permissionData)
	if ce.GinError(c, err) {
		return
	}
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
	p := c.Param("ID")
	id, err := strconv.Atoi(p)
	if ce.GinError(c, err) {
		return
	}

	permission, err := h.db.DeletePermission(id)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, permission)
}

type PermissionQuery struct {
	SubjectID   int    `gorm:"column:SubjectID"    json:"SubjectID"`
	ServiceName string `gorm:"column:PermissionServiceName"  json:"ServiceName"`
	Name        string `gorm:"column:PermissionName"          json:"Name"`
	Action      string `gorm:"column:PermissionAction"        json:"Action"`
}

// @Summary Permission 에 해당하는 objects 조회
// @Tags RBAC permission
// @Accept json
// @Produce json
// @Param data body PermissionQuery true "Object를 구하는데 필요한 permission 정보"
// @Success 200 {object} dblayer.PermissionAnswer "허용된 object list"
// @Router /rbac/permission/objects [post]
func (h *Handler) GetAllowedObjects(c *gin.Context) {

	var permissionKey PermissionQuery

	err := c.ShouldBindJSON(&permissionKey)
	if ce.GinError(c, err) {
		return
	}

	subjectID := permissionKey.SubjectID
	permissionServiceName := permissionKey.ServiceName
	permissionName := permissionKey.Name
	permissionAction := permissionKey.Action

	var permissionAnswer dblayer.PermissionAnswer
	permissionAnswer, err = h.db.GetAllowedObjects(
		subjectID,
		permissionServiceName,
		permissionName,
		permissionAction,
	)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, permissionAnswer)
}
