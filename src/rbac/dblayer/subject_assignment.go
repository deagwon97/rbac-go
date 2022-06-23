package dblayer

import (
	"rbac-go/rbac/models"
)

func (db *DBORM) GetSubjectAssignments() (
	items []models.SubjectAssignment, err error,
) {
	err = db.
		Find(&items).
		Error
	return
}

type SubjectAssignmentData struct {
	SubjectID int `gorm:"column:subject_id"    json:"subject_id"`
	RoleID    int `gorm:"column:role_id"       json:"role_id"`
}

func (db *DBORM) AddSubjectAssignment(
	itemData SubjectAssignmentData) (
	item models.SubjectAssignment, err error,
) {
	item.SubjectID =
		itemData.SubjectID
	item.RoleID =
		itemData.RoleID

	err = db.Create(&item).Error

	return item, err
}

func (db *DBORM) DeleteSubjectAssignment(
	itemData SubjectAssignmentData,
) (
	item models.SubjectAssignment,
	err error,
) {
	err = db.Raw(`
		SELECT * FROM subject_assignment 
		WHERE subject_id = ? AND role_id = ?`,
		itemData.SubjectID,
		itemData.RoleID,
	).First(&item).Error

	err = db.Delete(&item).Error

	return item, err
}
