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
	RoleID       int `gorm:"column:role_id"         json:"role_id"`
	PermissionID int `gorm:"column:permission_id"   json:"permission_id"`
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
		SELECT * FROM permission_assignment 
		WHERE permission_id = ? AND role_id = ?;`,
		itemData.PermissionID,
		itemData.RoleID,
	).First(&item).Error

	err = db.Delete(&item).Error

	return item, db.Delete(&item).Error
}
