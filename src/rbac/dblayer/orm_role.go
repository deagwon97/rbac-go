package dblayer

import (
	"errors"

	"rbac-go/common/paginate"
	"rbac-go/rbac/models"
)

// Role
func (db *DBORM) GetRoles() (
	roles []models.Role, err error,
) {
	err = db.
		Find(&roles).
		Error
	return
}

type RolesPage struct {
	Count        int           `json:"count"`
	NextPage     string        `json:"next"`
	PreviousPage string        `json:"previous"`
	Roles        []models.Role `json:"results"`
}

func (db *DBORM) GetRolesPage(
	page int, pageSize int, hostUrl string,
) (
	rolesPage RolesPage, err error,
) {

	var count int64
	db.Model(&models.Role{}).Count(&count)

	page, pageSize, nextPage, previousPage :=
		paginate.GetPageInfo(page, pageSize, hostUrl, count)
	rolesPage.Count = int(count)
	rolesPage.NextPage = nextPage
	rolesPage.PreviousPage = previousPage

	err = db.
		Select("id", "name", "description").
		Order("id desc").
		Scopes(paginate.Paginate(page, pageSize)).
		Find(&rolesPage.Roles).
		Error

	return rolesPage, err
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

type SubjectsOfRole struct {
	SubjectID int `json:"subject_id"`
}

type SubjectsOfRolePage struct {
	Count        int              `json:"count"`
	NextPage     string           `json:"next"`
	PreviousPage string           `json:"previous"`
	Items        []SubjectsOfRole `json:"results"`
}

func (db *DBORM) GetSubjectsOfRolePage(
	roleID int, page int, pageSize int, hostUrl string,
) (
	itemsPage SubjectsOfRolePage, err error,
) {
	var count int64
	err = db.Table("subject_assignment").
		Select("subject_id").
		Where("role_id = ?", roleID).
		Count(&count).
		Error

	page, pageSize, nextPage, previousPage :=
		paginate.GetPageInfo(page, pageSize, hostUrl, count)
	itemsPage.Count = int(count)
	itemsPage.NextPage = nextPage
	itemsPage.PreviousPage = previousPage

	err = db.Table("subject_assignment").
		Select("subject_id").
		Where("role_id = ?", roleID).
		Scopes(paginate.Paginate(page, pageSize)).
		Find(&itemsPage.Items).
		Error
	return itemsPage, err
}

type PermissionOfRole struct {
	PermissionID int    `gorm:"column:permission_id"   json:"permission_id"`
	ServiceName  string `gorm:"column:service_name"  json:"service_name"`
	Name         string `gorm:"column:name"          json:"name"`
	Action       string `gorm:"column:action"        json:"action"`
	Object       string `gorm:"column:object"        json:"object"`
}

type PermissionsOfRolePage struct {
	Count        int                `json:"count"`
	NextPage     string             `json:"next"`
	PreviousPage string             `json:"previous"`
	Items        []PermissionOfRole `json:"results"`
}

func (db *DBORM) GetPermissionsOfRolePage(
	roleID int, page int, pageSize int, hostUrl string,
) (
	itemsPage PermissionsOfRolePage, err error,
) {
	var count int64
	db.Model(&models.Permission{}).Count(&count)

	err = db.Table("permission_assignment").
		Select("permission_id",
			"service_name", "name",
			"action", "object").
		Where("role_id = ?", roleID).
		Order("permission_id desc").
		Joins(`left join 
		permission on 
		permission.id = permission_assignment.permission_id`).
		Count(&count).
		Error

	page, pageSize, nextPage, previousPage :=
		paginate.GetPageInfo(page, pageSize, hostUrl, count)
	itemsPage.Count = int(count)
	itemsPage.NextPage = nextPage
	itemsPage.PreviousPage = previousPage

	err = db.Table("permission_assignment").
		Select("role_id", "permission_id",
			"service_name", "name",
			"action", "object").
		Where("role_id = ?", roleID).
		Order("permission_id desc").
		Joins(`left join 
	permission on 
	permission.id = permission_assignment.permission_id`).
		Scopes(paginate.Paginate(page, pageSize)).
		Scan(&itemsPage.Items).
		Error
	return itemsPage, err
}
