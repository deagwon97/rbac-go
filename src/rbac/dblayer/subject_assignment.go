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
	SubjectID int `gorm:"column:SubjectID"    json:"SubjectID"`
	RoleID    int `gorm:"column:RoleID"       json:"RoleID"`
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
		SELECT * FROM SubjectAssignment 
		WHERE SubjectID = ? AND RoleID = ?`,
		itemData.SubjectID,
		itemData.RoleID,
	).First(&item).Error
	if err != nil {
		return
	}

	err = db.Delete(&item).Error
	if err != nil {
		return
	}

	return item, err
}
