package rest

import (
	"net/http"
	"rbac-go/rbac/dblayer"

	// "rbac-go/rbac/models"
	"rbac-go/database"

	"rbac-go/rbac/checker"

	"github.com/gin-gonic/gin"

	"strconv"
)

type Handler struct {
	db dblayer.DBLayer
}

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
	roles, err := h.db.GetRoles(page, pageSize, hostUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)
}

type HandlerInterface interface {
	GetRolesPage(c *gin.Context)
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
