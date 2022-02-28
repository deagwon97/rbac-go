package dblayer

import (
	"database/sql"
	"errors"
	"fmt"

	"rbac-go/common/paginate"
	"rbac-go/rbac/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBORM struct {
	*gorm.DB
}

// DBORM의 생성자
func NewORM(dbengine string, dsn string) (*DBORM, error) {
	sqlDB, err := sql.Open(dbengine, dsn)
	if err != nil {
		return &DBORM{
			DB: nil,
		}, err
	}
	gormDB, err := gorm.Open(
		mysql.New(mysql.Config{Conn: sqlDB}),
		&gorm.Config{},
	)
	return &DBORM{
		DB: gormDB,
	}, err
}

//  Permission
type PermissionObject struct {
	PermissionObject string `gorm:"column:permission_object" json:"permission_object"`
}

func (db *DBORM) GetObjects(
	subjectID int,
	permissionServiceName string,
	permissionName string,
	permissionAction string,
) (
	objects []string,
	err error,
) {
	var permissionObject []PermissionObject

	query := fmt.Sprintf(`
	SELECT pa.permission_object
	FROM permission_assignment pa
	INNER JOIN permission r
		ON r.id = pa.permission_id
	INNER JOIN subject_assignment s
		ON r.id = s.permission_id
	WHERE s.id = %d 
	AND pa.permission_name ='%s' 
	AND pa.permission_action = '%s' 
	`, subjectID, permissionName, permissionAction)

	err = db.Raw(query).Scan(&permissionObject).Error
	if err != nil {
		return objects, err
	}

	for _, item := range permissionObject {
		objects = append(objects, item.PermissionObject)
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
		Select("id", "subject_name",
			"name", "action", "object").
		Order("id desc").
		Scopes(paginate.Paginate(page, pageSize)).
		Find(&permissionPage.Permissions).
		Error

	return permissionPage, err
}

type PermissionData struct {
	ServiceName string `gorm:"column:subject_name"  json:"subject_name"`
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
	err = db.Model(&permission).Updates(permission).Error

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

// Role
func (db *DBORM) GetRoles() (
	roles []models.Role, err error,
) {
	err = db.
		Find(&roles).
		Error
	return
}

type RolePage struct {
	Count        int           `json:"count"`
	NextPage     string        `json:"next"`
	PreviousPage string        `json:"previous"`
	Roles        []models.Role `json:"results"`
}

func (db *DBORM) GetRolesPage(
	page int, pageSize int, hostUrl string,
) (
	rolePage RolePage, err error,
) {

	var count int64
	db.Model(&models.Role{}).Count(&count)

	page, pageSize, nextPage, previousPage :=
		paginate.GetPageInfo(page, pageSize, hostUrl, count)
	rolePage.Count = int(count)
	rolePage.NextPage = nextPage
	rolePage.PreviousPage = previousPage

	err = db.
		Select("id", "name", "description").
		Order("id desc").
		Scopes(paginate.Paginate(page, pageSize)).
		Find(&rolePage.Roles).
		Error

	return rolePage, err
}

type RoleData struct {
	Name        string `gorm:"column:name"           json:"name"`
	Description string `gorm:"column:description"    json:"description"`
}

func (db *DBORM) AddRole(roleData RoleData) (
	role models.Role, err error,
) {
	role.Name = roleData.Name
	role.Description = roleData.Description
	err = db.Create(&role).Error
	return role, err
}

func (db *DBORM) UpdateRole(
	id int,
	roleData RoleData,
) (
	role models.Role,
	err error,
) {
	role.ID = id
	role.Name = roleData.Name
	role.Description = roleData.Description

	var count int64
	db.Model(&models.Role{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		return role, errors.New("item dosen't exist")
	}

	err = db.Model(role).Updates(role).Error
	db.Save(&role)

	db.Where("id = ?", id).First(&role)
	return role, err
}

func (db *DBORM) DeleteRole(
	id int,
) (
	role models.Role,
	err error,
) {
	db.Where("id = ?", id).First(&role)
	return role, db.Delete(&role).Error
}
