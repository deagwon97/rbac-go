package rest

import (
	"net/http"
	ce "rbac-go/common/error"
	"rbac-go/rbac/dblayer"

	"github.com/gin-gonic/gin"

	"strconv"
)

// @Summary SubjectAssignment 생성
// @Tags RBAC subjectAssignment
// @Accept json
// @Produce json
// @Param data body dblayer.SubjectAssignmentData true "Data"
// @Success 200 {object} models.SubjectAssignment
// @Router /rbac/subjectAssignment [post]
func (h *Handler) AddSubjectAssignment(c *gin.Context) {
	var itemData dblayer.SubjectAssignmentData

	err := c.ShouldBindJSON(&itemData)
	ce.GinError(c, err)

	item, err := h.db.AddSubjectAssignment(itemData)
	ce.GinError(c, err)
	c.JSON(http.StatusOK, item)
}

// @Summary SubjectAssignment Update
// @Tags RBAC subjectAssignment
// @Accept json
// @Produce json
// @Param id path int  true  "SubjectAssignment ID"
// @Param data body dblayer.SubjectAssignmentData true "Update에 사용할 Data"
// @Success 200 {object} models.SubjectAssignment "수정된 SubjectAssignment 데이터"
// @Router /rbac/subjectAssignment/{id} [patch]
func (h *Handler) UpdateSubjectAssignment(c *gin.Context) {

	defer gin.Recovery()

	p := c.Param("id")
	id, err := strconv.Atoi(p)
	ce.GinError(c, err)
	var itemData dblayer.SubjectAssignmentData

	err = c.ShouldBindJSON(&itemData)
	ce.GinError(c, err)

	item, err := h.db.UpdateSubjectAssignment(id, itemData)
	ce.GinError(c, err)
	c.JSON(http.StatusOK, item)
}

// @Summary SubjectAssignment 삭제
// @Tags RBAC subjectAssignment
// @Accept json
// @Produce json
// @Param id path int true  "SubjectAssignment ID"
// @Success 200 {object} models.SubjectAssignment "삭제된 SubjectAssignment 데이터"
// @Router /rbac/subjectAssignment/{id} [delete]
func (h *Handler) DeleteSubjectAssignment(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	ce.GinError(c, err)

	item, err := h.db.DeleteSubjectAssignment(id)
	ce.GinError(c, err)
	c.JSON(http.StatusOK, item)
}
