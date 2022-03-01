package dblayer

import (
	"errors"

	"rbac-go/rbac/models"
)

// SubjectAssignment
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

func (db *DBORM) UpdateSubjectAssignment(
	id int,
	itemData SubjectAssignmentData,
) (
	item models.SubjectAssignment,
	err error,
) {
	item.ID = id
	item.SubjectID =
		itemData.SubjectID
	item.RoleID =
		itemData.RoleID

	var count int64
	db.Model(item).Where("id = ?", id).Count(&count)
	if count == 0 {
		return item, errors.New("item dosen't exist")
	}

	err = db.Model(item).Updates(item).Error
	db.Save(&item)

	db.Where("id = ?", id).First(&item)
	return item, err
}

func (db *DBORM) DeleteSubjectAssignment(
	id int,
) (
	item models.SubjectAssignment,
	err error,
) {
	db.Where("id = ?", id).First(&item)
	return item, db.Delete(&item).Error
}
