package rest

import (
	"net/http"
	"rbac/content/dblayer"
	"rbac/content/models"
	"rbac/database"

	"github.com/gin-gonic/gin"

	"strconv"
)

type Handler struct {
	db dblayer.DBLayer
}

// @BasePath /

// Go API godoc
// @Summary Content 목록 조회
// @Schemes
// @Description Content 목록 조회
// @Tags Content
// @Accept json
// @Produce json
// @Param page query int  false  "Page Number"
// @Param pageSize query int  false  "Page Size"
// @Success 200 {object} models.ContentList
// @Router /content/list [get]
func (h *Handler) GetContents(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	hostName := c.Request.Host + "/content/list"
	scheme := "http://"
	if c.Request.TLS != nil {
		scheme = "https://"
	}
	hostUrl := scheme + hostName

	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "dsn 오류"})
		return
	}
	contents, err := h.db.GetAllContents(page, pageSize, hostUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contents)
}

// Go API godoc
// @Summary Content 상세 조회
// @Schemes
// @Description Content 상세 조회
// @Tags Content
// @Accept json
// @Produce json
// @Param id path int true  "Content id"
// @Success 200 {object} models.Content
// @Router /content/{id} [get]
func (h *Handler) GetContent(c *gin.Context) {

	p := c.Param("id")
	id, err := strconv.Atoi(p)

	content, err := h.db.GetContent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, content)
}

// Go API godoc
// @Summary Content 생성
// @Schemes
// @Tags Content
// @Accept json
// @Produce json
// @Param data body models.ContentData true  "Content Data"
// @Success 200 {object} models.Content
// @Router /content [post]
func (h *Handler) AddContent(c *gin.Context) {

	var contentData models.ContentData

	err := c.ShouldBindJSON(&contentData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	content, err := h.db.AddContent(contentData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, content)
}

// Go API godoc
// @Summary Content 수정
// @Schemes
// @Tags Content
// @Accept json
// @Produce json
// @Param id path int true  "Content id"
// @Param data body models.ContentData true  "Content Data"
// @Success 200 {object} models.Content
// @Router /content/{id} [patch]
func (h *Handler) UpdateContent(c *gin.Context) {

	p := c.Param("id")
	id, err := strconv.Atoi(p)
	var contentData models.ContentData

	err = c.ShouldBindJSON(&contentData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	content, err := h.db.UpdateContent(id, contentData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, content)
}

// Go API godoc
// @Summary Content 삭제
// @Schemes
// @Tags Content
// @Accept json
// @Produce json
// @Param id path int true  "Content id"
// @Success 200 {object} models.Content
// @Router /content/{id} [delete]
func (h *Handler) DeleteContent(c *gin.Context) {

	p := c.Param("id")
	id, err := strconv.Atoi(p)

	content, err := h.db.DeleteContent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, content)
}

type HandlerInterface interface {
	GetContents(c *gin.Context)
	GetContent(c *gin.Context)
	AddContent(c *gin.Context)
	UpdateContent(c *gin.Context)
	DeleteContent(c *gin.Context)
}

// HandlerInterface의 생성자
func NewHandler() (HandlerInterface, error) {
	dsn := database.DataSource
	// DBORM 초기화
	db, err := dblayer.NewORM("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}
