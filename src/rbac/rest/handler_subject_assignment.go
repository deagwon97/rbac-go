package rest

import (
	"net/http"
	ce "rbac-go/common/error"
	"rbac-go/rbac/dblayer"
	"strconv"

	"github.com/gin-gonic/gin"
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

// @Summary SubjectAssignment 삭제
// @Tags RBAC subjectAssignment
// @Accept json
// @Produce json
// @Param data body dblayer.SubjectAssignmentData true "Data"
// @Success 200 {object} models.SubjectAssignment "삭제된 SubjectAssignment 데이터"
// @Router /rbac/subject-assignment [delete]
func (h *Handler) DeleteSubjectAssignment(c *gin.Context) {
	var subjectAssignmentData dblayer.SubjectAssignmentData

	subjectID, _ := strconv.Atoi(c.Query("subjectID"))
	roleID, _ := strconv.Atoi(c.Query("roleID"))

	subjectAssignmentData.SubjectID = subjectID
	subjectAssignmentData.RoleID = roleID

	item, err := h.db.DeleteSubjectAssignment(subjectAssignmentData)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}
