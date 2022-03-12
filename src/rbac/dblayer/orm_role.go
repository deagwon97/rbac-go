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
	RoleID        int   `gorm:"column:role_id"       json:"role_id"`
	SubjectIDList []int `json:"subject_id_list"`
}

type SubjectStatus struct {
	SubjectID int  `gorm:"column:subject_id"   json:"subject_id"`
	IsAllowed bool `json:"is_allowed"`
}

type SubjectStatusOfRole struct {
	RoleID int             `gorm:"column:role_id"       json:"role_id"`
	List   []SubjectStatus `json:"list"`
}

func (db *DBORM) CheckSubjectIsAllowed(
	subjectsOfRole SubjectsOfRole,
) (
	subjectStatusOfRole SubjectStatusOfRole, err error,
) {
	roleID := subjectsOfRole.RoleID
	var allowedSubjectIDList []int
	err = db.Table("subject_assignment").
		Select("subject_id").
		Where("subject_id IN ?", subjectsOfRole.SubjectIDList).
		Where("role_id = ?", roleID).
		Scan(&allowedSubjectIDList).
		Error

	var subjectStatus SubjectStatus
	subjectStatusOfRole.RoleID = roleID
	for _, subjectID := range subjectsOfRole.SubjectIDList {
		subjectStatus.SubjectID = subjectID
		if utils.IsIn(subjectID, allowedSubjectIDList) == true {
			subjectStatus.IsAllowed = true
		} else {
			subjectStatus.IsAllowed = false
		}
		subjectStatusOfRole.List = append(subjectStatusOfRole.List, subjectStatus)
	}
	return subjectStatusOfRole, err
}

type PermissionOfRole struct {
	RoleID           int   `gorm:"column:role_id"       json:"role_id"`
	PermissionIDList []int `json:"permission_id_list"`
}

type PermissionStatus struct {
	PermissionID int  `gorm:"column:permission_id"   json:"permission_id"`
	IsAllowed    bool `json:"is_allowed"`
}

type PermissionStatusOfRole struct {
	RoleID int                `gorm:"column:role_id"       json:"role_id"`
	List   []PermissionStatus `json:"list"`
}

func (db *DBORM) CheckPermissionIsAllowed(
	permissionOfRole PermissionOfRole,
) (
	permissionStatusOfRole PermissionStatusOfRole, err error,
) {
	roleID := permissionOfRole.RoleID
	var allowedPermissionIDList []int

	err = db.Table("permission_assignment").
		Select("permission_id").
		Where("permission_id IN ?", permissionOfRole.PermissionIDList).
		Where("role_id = ?", roleID).
		Scan(&allowedPermissionIDList).
		Error

	var permissionStatus PermissionStatus
	permissionStatusOfRole.RoleID = roleID
	for _, permissionID := range permissionOfRole.PermissionIDList {
		permissionStatus.PermissionID = permissionID
		if utils.IsIn(permissionID, allowedPermissionIDList) == true {
			permissionStatus.IsAllowed = true
		} else {
			permissionStatus.IsAllowed = false
		}
		permissionStatusOfRole.List = append(permissionStatusOfRole.List, permissionStatus)
	}
	return permissionStatusOfRole, err
}
