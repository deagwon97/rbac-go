package dblayer

import (
	"errors"

	"rbac-go/common/utils"

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
	Count        int           `json:"Count"`
	NextPage     string        `json:"NextPage"`
	PreviousPage string        `json:"PreviousPage"`
	Roles        []models.Role `json:"Results"`
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
		Select("ID", "Name", "Description").
		Order("ID desc").
		Scopes(paginate.Paginate(page, pageSize)).
		Find(&rolesPage.Roles).
		Error

	return rolesPage, err
}

type RoleData struct {
	Name        string `gorm:"column:Name"           json:"Name"`
	Description string `gorm:"column:Description"    json:"Description"`
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
	db.Model(&models.Role{}).Where("ID = ?", id).Count(&count)
	if count == 0 {
		return role, errors.New("item dosen't exist")
	}

	err = db.Model(role).Updates(role).Error
	db.Save(&role)

	db.Where("ID = ?", id).First(&role)
	return role, err
}

func (db *DBORM) DeleteRole(
	id int,
) (
	role models.Role,
	err error,
) {
	db.Where("ID = ?", id).First(&role)
	return role, db.Delete(&role).Error
}

type SubjectsOfRole struct {
	RoleID        int   `gorm:"column:RoleID"       json:"RoleID"`
	SubjectIDList []int `json:"SubjectIDList"`
}

type SubjectStatus struct {
	SubjectID int  `gorm:"column:SubjectID"   json:"SubjectID"`
	IsAllowed bool `json:"IsAllowed"`
}

type SubjectStatusOfRole struct {
	RoleID int             `gorm:"column:RoleID"       json:"RoleID"`
	List   []SubjectStatus `json:"List"`
}

func (db *DBORM) CheckSubjectIsAllowed(
	subjectsOfRole SubjectsOfRole,
) (
	subjectStatusOfRole SubjectStatusOfRole, err error,
) {
	roleID := subjectsOfRole.RoleID
	var allowedSubjectIDList []int
	err = db.Table("SubjectAssignment").
		Select("SubjectID").
		Where("SubjectID IN ?", subjectsOfRole.SubjectIDList).
		Where("RoleID = ?", roleID).
		Scan(&allowedSubjectIDList).
		Error

	var subjectStatus SubjectStatus
	subjectStatusOfRole.RoleID = roleID
	for _, subjectID := range subjectsOfRole.SubjectIDList {
		subjectStatus.SubjectID = subjectID
		if utils.IsIn(subjectID, allowedSubjectIDList) {
			subjectStatus.IsAllowed = true
		} else {
			subjectStatus.IsAllowed = false
		}
		subjectStatusOfRole.List = append(subjectStatusOfRole.List, subjectStatus)
	}
	return subjectStatusOfRole, err
}

type PermissionOfRole struct {
	RoleID           int   `gorm:"column:RoleID"       json:"RoleID"`
	PermissionIDList []int `json:"PermissionIDList"`
}

type PermissionStatus struct {
	PermissionID int  `gorm:"column:PermissionID"   json:"PermissionID"`
	IsAllowed    bool `json:"IsAllowed"`
}

type PermissionStatusOfRole struct {
	RoleID int                `gorm:"column:RoleID"       json:"RoleID"`
	List   []PermissionStatus `json:"List"`
}

func (db *DBORM) CheckPermissionIsAllowed(
	permissionOfRole PermissionOfRole,
) (
	permissionStatusOfRole PermissionStatusOfRole, err error,
) {
	roleID := permissionOfRole.RoleID
	var allowedPermissionIDList []int

	err = db.Table("PermissionAssignment").
		Select("PermissionID").
		Where("PermissionID IN ?", permissionOfRole.PermissionIDList).
		Where("RoleID = ?", roleID).
		Scan(&allowedPermissionIDList).
		Error

	var permissionStatus PermissionStatus
	permissionStatusOfRole.RoleID = roleID
	for _, permissionID := range permissionOfRole.PermissionIDList {
		permissionStatus.PermissionID = permissionID
		if utils.IsIn(permissionID, allowedPermissionIDList) {
			permissionStatus.IsAllowed = true
		} else {
			permissionStatus.IsAllowed = false
		}
		permissionStatusOfRole.List = append(permissionStatusOfRole.List, permissionStatus)
	}
	return permissionStatusOfRole, err
}
