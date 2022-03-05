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

type PermissionAnswer struct {
	Objects   []string `json:"objects"`
	IsAllowed bool     `json:"is_allowed"`
}

func (db *DBORM) GetAllowedObjects(
	subjectID int,
	permissionServiceName string,
	permissionName string,
	permissionAction string,
) (
	permissionAnswer PermissionAnswer,
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

	var permissionObjects []PermissionObject
	err = db.Raw(query).Scan(&permissionObjects).Error
	if err != nil {
		return
	}

	if len(permissionObjects) > 0 {
		permissionAnswer.IsAllowed = true
		for _, item := range permissionObjects {
			permissionAnswer.Objects =
				append(permissionAnswer.Objects, item.Object)
		}
	}
	return permissionAnswer, err
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

type PermissionSet struct {
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
	Objects []string `json:"objects"`
}

type PermissionSetData struct {
	ServiceName    string          `json:"service_name"`
	PermissionSets []PermissionSet `json:"permission_sets"`
}

type TempPermission struct {
	models.Permission
}

func (TempPermission) TableName() string {
	return "temp_permission"
}

func (db *DBORM) AddPermissionSets(permissionSetData PermissionSetData) (
	permissions []TempPermission, err error,
) {
	var permission TempPermission
	permission.ServiceName = permissionSetData.ServiceName

	if len(permissionSetData.PermissionSets) > 0 {
		for _, permissionSet := range permissionSetData.PermissionSets {

			permission.Name = permissionSet.Name

			if len(permissionSet.Actions) > 0 {
				for _, action := range permissionSet.Actions {

					permission.Action = action

					if len(permissionSet.Objects) > 0 {
						for _, object := range permissionSet.Objects {

							permission.Object = object
							permissions = append(permissions, permission)
						}
					} else {
						permissions = append(permissions, permission)
					}
				}
			}
		}

		tx := db.Begin()

		qeury := fmt.Sprintf(`
		CREATE TEMPORARY TABLE temp_permission( 
			service_name varchar(64) DEFAULT NULL,
			name varchar(64) DEFAULT NULL,
			action varchar(64) DEFAULT NULL,
			object varchar(64) DEFAULT NULL,
			UNIQUE KEY service_name_name_action_object (
			service_name, name, action, object)
		  );`)

		if err = tx.Exec(qeury).Error; err != nil {
			fmt.Println(err)
			return
		}

		if err = tx.CreateInBatches(&permissions, 100).Error; err != nil {
			fmt.Println(err)
			return
		}

		if err = tx.Exec(`
			INSERT IGNORE INTO permission(service_name, name, action, object) 
			SELECT * FROM temp_permission;
				`).Error; err != nil {
			fmt.Println(err)
			return
		}

		if err = tx.Exec(`
			DELETE FROM p USING permission AS p
			WHERE NOT EXISTS(
				SELECT * FROM temp_permission tp
				WHERE
				p.service_name = tp.service_name AND 
				p.name = tp.name AND
				p.action = tp.action AND
				p.object = tp.object
			);
				`).Error; err != nil {
			fmt.Println(err)
			return
		}

		if err = tx.Exec(`
			DROP TABLE temp_permission;
				`).Error; err != nil {
			fmt.Println(err)
			return
		}

		tx.Commit()
		return
	}
	return permissions, err
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
