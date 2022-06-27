package dblayer

import (
	"rbac-go/rbac/models"
)

func (db *DBORM) GetPermissionAssignments() (
	items []models.PermissionAssignment, err error,
) {
	err = db.
		Find(&items).
		Error
	return
}

type PermissionAssignmentData struct {
	RoleID       int `gorm:"column:RoleID"         json:"RoleID"`
	PermissionID int `gorm:"column:PermissionID"   json:"PermissionID"`
}

func (db *DBORM) AddPermissionAssignment(
	itemData PermissionAssignmentData) (
	item models.PermissionAssignment, err error,
) {
	item.RoleID =
		itemData.RoleID
	item.PermissionID =
		itemData.PermissionID
	err = db.Create(&item).Error

	return item, err
}

func (db *DBORM) DeletePermissionAssignment(
	itemData PermissionAssignmentData,
) (
	item models.PermissionAssignment,
	err error,
) {
	err = db.Raw(`
		SELECT * FROM PermissionAssignment 
		WHERE PermissionID = ? AND RoleID = ?;`,
		itemData.PermissionID,
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
