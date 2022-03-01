package dblayer

import (
	"errors"
	"fmt"

	"rbac-go/common/paginate"
	"rbac-go/rbac/models"
)

type PermissionObject struct {
	Object string `gorm:"column:object" json:"object"`
}

func (db *DBORM) GetAllowedObjects(
	subjectID int,
	permissionServiceName string,
	permissionName string,
	permissionAction string,
) (
	objects []string,
	err error,
) {

	query := fmt.Sprintf(`
	SELECT p.object
	FROM role as r
	INNER JOIN permission_assignment as pa
		ON r.id = pa.role_id
	INNER JOIN permission as p
		ON pa.permission_id = p.id
	INNER JOIN subject_assignment as sa
		ON r.id = sa.role_id
	WHERE sa.subject_id = %d 
	AND p.service_name ='%s' 
	AND p.name ='%s' 
	AND p.action = '%s' 
	`, subjectID, permissionServiceName, permissionName, permissionAction)

	var permissionObject []PermissionObject
	err = db.Raw(query).Scan(&permissionObject).Error
	if err != nil {
		return
	}

	for _, item := range permissionObject {
		objects = append(objects, item.Object)
	}
	return objects, err
}

func (db *DBORM) GetPermissions() (
	permissions []models.Permission, err error,
) {
	err = db.
		Find(&permissions).
		Error
	return
}

type PermissionsPage struct {
	Count        int                 `json:"count"`
	NextPage     string              `json:"next"`
	PreviousPage string              `json:"previous"`
	Permissions  []models.Permission `json:"results"`
}

func (db *DBORM) GetPermissionsPage(
	page int, pageSize int, hostUrl string,
) (
	permissionPage PermissionsPage, err error,
) {

	var count int64
	db.Model(&models.Permission{}).Count(&count)

	page, pageSize, nextPage, previousPage :=
		paginate.GetPageInfo(page, pageSize, hostUrl, count)
	permissionPage.Count = int(count)
	permissionPage.NextPage = nextPage
	permissionPage.PreviousPage = previousPage

	err = db.
		Select("id", "service_name",
			"name", "action", "object").
		Order("id desc").
		Scopes(paginate.Paginate(page, pageSize)).
		Find(&permissionPage.Permissions).
		Error

	return permissionPage, err
}

type PermissionData struct {
	ServiceName string `gorm:"column:service_name"  json:"service_name"`
	Name        string `gorm:"column:name"          json:"name"`
	Action      string `gorm:"column:action"        json:"action"`
	Object      string `gorm:"column:object"        json:"object"`
}

func (db *DBORM) AddPermission(permissionData PermissionData) (
	permission models.Permission, err error,
) {
	permission.ServiceName = permissionData.ServiceName
	permission.Name = permissionData.Name
	permission.Action = permissionData.Action
	permission.Object = permissionData.Object
	err = db.Create(&permission).Error
	return permission, err
}

func (db *DBORM) UpdatePermission(
	id int,
	permissionData PermissionData,
) (
	permission models.Permission,
	err error,
) {
	permission.ID = id
	permission.ServiceName = permissionData.ServiceName
	permission.Name = permissionData.Name
	permission.Action = permissionData.Action
	permission.Object = permissionData.Object

	var count int64
	db.Model(&models.Permission{}).
		Where("id = ?", id).
		Count(&count)
	if count == 0 {
		return permission,
			errors.New("item dosen't exist")
	}
	err = db.Model(permission).
		Updates(permission).Error

	db.Where("id = ?", id).First(&permission)
	return permission, err
}

func (db *DBORM) DeletePermission(
	id int,
) (
	permission models.Permission,
	err error,
) {
	db.Where("id = ?", id).First(&permission)
	return permission, db.Delete(&permission).Error
}
