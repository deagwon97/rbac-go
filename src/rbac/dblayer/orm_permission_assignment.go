package dblayer

import (
	"errors"

	"rbac-go/rbac/models"
)

// PermissionAssignment
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

func (db *DBORM) UpdatePermissionAssignment(
	id int,
	itemData PermissionAssignmentData,
) (
	item models.PermissionAssignment,
	err error,
) {
	item.ID = id
	item.RoleID =
		itemData.RoleID
	item.PermissionID =
		itemData.PermissionID

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

func (db *DBORM) DeletePermissionAssignment(
	id int,
) (
	item models.PermissionAssignment,
	err error,
) {
	db.Where("id = ?", id).First(&item)
	return item, db.Delete(&item).Error
}
