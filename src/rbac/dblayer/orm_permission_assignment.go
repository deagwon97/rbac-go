package dblayer

import (
	"errors"

	"rbac-go/common/paginate"
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

type PermissionAssignmentsPage struct {
	Count        int                           `json:"count"`
	NextPage     string                        `json:"next"`
	PreviousPage string                        `json:"previous"`
	Items        []models.PermissionAssignment `json:"results"`
}

func (db *DBORM) GetPermissionAssignmentsPage(
	page int, pageSize int, hostUrl string,
) (
	itemsPage PermissionAssignmentsPage, err error,
) {

	var count int64
	db.Model(&models.PermissionAssignment{}).Count(&count)

	page, pageSize, nextPage, previousPage :=
		paginate.GetPageInfo(page, pageSize, hostUrl, count)
	itemsPage.Count = int(count)
	itemsPage.NextPage = nextPage
	itemsPage.PreviousPage = previousPage

	err = db.
		Select("id", "role_id", "permission_id").
		Order("id desc").
		Scopes(paginate.Paginate(page, pageSize)).
		Find(&itemsPage.Items).
		Error

	return itemsPage, err
}

type PermissionAssignmentData struct {
	Name         string `gorm:"column:name"           json:"name"`
	RoleID       int    `gorm:"column:role_id"         json:"role_id"`
	PermissionID string `gorm:"column:permission_id"   json:"spermission_id"`
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
	db.Model(&models.PermissionAssignment{}).Where("id = ?", id).Count(&count)
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
