package rest

import (
	"net/http"
	ce "rbac-go/common/error"
	"rbac-go/common/paginate"
	"rbac-go/rbac/dblayer"

	"rbac-go/database"

	"rbac-go/rbac/checker"

	"github.com/gin-gonic/gin"

	"strconv"
)

type Handler struct {
	db dblayer.DBLayer
}

type HandlerInterface interface {
	GetRolesPage(c *gin.Context)
	AddRole(c *gin.Context)
	UpdateRole(c *gin.Context)
	DeleteRole(c *gin.Context)
}

// HandlerInterface의 생성자
func NewHandler() (HandlerInterface, error) {

	// RBAC 초기화
	checker.NewRBAC()

	// DBORM 초기화
	dsn := database.DataSource
	db, err := dblayer.NewORM("mysql", dsn)
	ce.PanicIfError(err)
	return &Handler{
		db: db,
	}, nil
}

// @Summary Role 목록 조회
// @Tags RBAC
// @Accept json
// @Produce json
// @Param page query int false  "Page Number"
// @Param pageSize query int false  "Page Size"
// @Success 200 {object} dblayer.RolePage
// @Router /rbac/role/list [get]
func (h *Handler) GetRolesPage(c *gin.Context) {
	page, pageSize, hostUrl := paginate.ParsePageUrl(c)
	roles, err := h.db.GetRolesPage(page, pageSize, hostUrl)
	ce.GinError(c, err)
	c.JSON(http.StatusOK, roles)
}

// @Summary Role 생성
// @Tags RBAC
// @Accept json
// @Produce json
// @Param data body dblayer.RoleData true "Data"
// @Success 200 {object} models.Role
// @Router /rbac/role [post]
func (h *Handler) AddRole(c *gin.Context) {
	var roleData dblayer.RoleData

	err := c.ShouldBindJSON(&roleData)
	ce.GinError(c, err)

	role, err := h.db.AddRole(roleData)
	ce.GinError(c, err)
	c.JSON(http.StatusOK, role)
}

// @Summary Role Update
// @Tags RBAC
// @Accept json
// @Produce json
// @Param id path int  true  "Role ID"
// @Param data body dblayer.RoleData true "Update에 사용할 Data"
// @Success 200 {object} models.Role "수정된 Role 데이터"
// @Router /rbac/role/{id} [patch]
func (h *Handler) UpdateRole(c *gin.Context) {

	p := c.Param("id")
	id, err := strconv.Atoi(p)
	ce.GinError(c, err)
	var roleData dblayer.RoleData

	err = c.ShouldBindJSON(&roleData)
	ce.GinError(c, err)

	role, err := h.db.UpdateRole(id, roleData)
	ce.GinError(c, err)
	c.JSON(http.StatusOK, role)
}

// @Summary Role 삭제
// @Tags RBAC
// @Accept json
// @Produce json
// @Param id path int true  "Role ID"
// @Success 200 {object} models.Role "삭제된 Role 데이터"
// @Router /rbac/role/{id} [delete]
func (h *Handler) DeleteRole(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	ce.GinError(c, err)

	role, err := h.db.DeleteRole(id)
	ce.GinError(c, err)
	c.JSON(http.StatusOK, role)
}
