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
	// err = db.Model(models.Role{}).Create(roleData).Error

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
