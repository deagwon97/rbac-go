package rest

import (
	"net/http"
	"rbac-go/rbac/dblayer"

	"rbac-go/database"
	"rbac-go/rbac/models"

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
	if err != nil {
		return nil, err
	}
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
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	hostName := c.Request.Host + "/role/list"
	scheme := "http://"
	if c.Request.TLS != nil {
		scheme = "https://"
	}
	hostUrl := scheme + hostName

	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "dsn 오류"})
		return
	}
	var roles dblayer.RolePage
	var err error
	roles, err = h.db.GetRolesPage(page, pageSize, hostUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
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
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	role, err := h.db.AddRole(roleData)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}
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
	var roleData dblayer.RoleData

	err = c.ShouldBindJSON(&roleData)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	var role models.Role
	role, err = h.db.UpdateRole(id, roleData)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
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

	role, err := h.db.DeleteRole(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, role)
}
