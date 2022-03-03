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
// @Router /rbac/subject-assignment [post]
func (h *Handler) AddSubjectAssignment(c *gin.Context) {
	var itemData dblayer.SubjectAssignmentData

	err := c.ShouldBindJSON(&itemData)
	if ce.GinError(c, err) {
		return
	}

	item, err := h.db.AddSubjectAssignment(itemData)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

// @Summary SubjectAssignment Update
// @Tags RBAC subjectAssignment
// @Accept json
// @Produce json
// @Param id path int  true  "SubjectAssignment ID"
// @Param data body dblayer.SubjectAssignmentData true "Update에 사용할 Data"
// @Success 200 {object} models.SubjectAssignment "수정된 SubjectAssignment 데이터"
// @Router /rbac/subject-assignment/{id} [patch]
func (h *Handler) UpdateSubjectAssignment(c *gin.Context) {

	defer gin.Recovery()

	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if ce.GinError(c, err) {
		return
	}
	var itemData dblayer.SubjectAssignmentData

	err = c.ShouldBindJSON(&itemData)
	if ce.GinError(c, err) {
		return
	}

	item, err := h.db.UpdateSubjectAssignment(id, itemData)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

// @Summary SubjectAssignment 삭제
// @Tags RBAC subjectAssignment
// @Accept json
// @Produce json
// @Param id path int true  "SubjectAssignment ID"
// @Success 200 {object} models.SubjectAssignment "삭제된 SubjectAssignment 데이터"
// @Router /rbac/subject-assignment/{id} [delete]
func (h *Handler) DeleteSubjectAssignment(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if ce.GinError(c, err) {
		return
	}

	item, err := h.db.DeleteSubjectAssignment(id)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}
