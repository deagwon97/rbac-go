package dblayer

import (
	"errors"
	"fmt"

	"rbac-go/common/paginate"
	"rbac-go/rbac/models"
)

type PermissionObject struct {
	Object string `gorm:"column:Object" json:"Object"`
}

type PermissionAnswer struct {
	Objects   []string `json:"Objects"`
	IsAllowed bool     `json:"IsAllowed"`
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
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
	SELECT p.Object
	FROM Role as r
	INNER JOIN PermissionAssignment as pa
		ON r.ID = pa.RoleID
	INNER JOIN Permission as p
		ON pa.PermissionID = p.ID
	INNER JOIN SubjectAssignment as sa
		ON r.ID = sa.RoleID
	WHERE sa.SubjectID = %d 
	AND p.ServiceName ='%s' 
	AND p.Name ='%s' 
	AND p.Action = '%s' 
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
	permissionAnswer.Objects = removeDuplicateStr(permissionAnswer.Objects)
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
	Count        int                 `json:"Count"`
	NextPage     string              `json:"NextPage"`
	PreviousPage string              `json:"PreviousPage"`
	Permissions  []models.Permission `json:"Results"`
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
		Select("ID", "ServiceName",
			"Name", "Action", "Object").
		Order("ID desc").
		Scopes(paginate.Paginate(page, pageSize)).
		Find(&permissionPage.Permissions).
		Error

	return permissionPage, err
}

type PermissionsStatus struct {
	models.Permission
	IsAllowed bool `json:"IsAllowed"`
}

type PermissionsStatusPage struct {
	Count        int                 `json:"Count"`
	NextPage     string              `json:"NextPage"`
	PreviousPage string              `json:"PreviousPage"`
	List         []PermissionsStatus `json:"Results"`
}

func (db *DBORM) GetPermissionsStatusPage(
	roleID int, page int, pageSize int, hostUrl string,
) (
	permissionPage PermissionsStatusPage, err error,
) {

	var count int64
	db.Model(&models.Permission{}).Count(&count)

	page, pageSize, nextPage, previousPage :=
		paginate.GetPageInfo(page, pageSize, hostUrl, count)
	permissionPage.Count = int(count)
	permissionPage.NextPage = nextPage
	permissionPage.PreviousPage = previousPage

	err = db.
		Select("ID", "ServiceName",
			"Name", "Action", "Object").
		Order("ID desc").
		Scopes(paginate.Paginate(page, pageSize)).
		Find(&permissionPage.List).
		Error
	if err != nil {
		return permissionPage, err
	}

	var permissionIDList []int
	for _, permission := range permissionPage.List {
		permissionIDList = append(permissionIDList, permission.ID)
	}
	permissionOfRole := PermissionOfRole{
		RoleID:           roleID,
		PermissionIDList: permissionIDList,
	}
	permissionStatusOfRole, err := db.CheckPermissionIsAllowed(permissionOfRole)
	for idx := 0; idx < len(permissionPage.List); idx++ {
		permissionPage.List[idx].IsAllowed = permissionStatusOfRole.List[idx].IsAllowed
	}

	return permissionPage, err
}

type PermissionData struct {
	ServiceName string `gorm:"column:ServiceName"  json:"ServiceName"`
	Name        string `gorm:"column:Name"          json:"Name"`
	Action      string `gorm:"column:Action"        json:"Action"`
	Object      string `gorm:"column:Object"        json:"Object"`
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
	Name    string   `json:"Name"`
	Actions []string `json:"Actions"`
	Objects []string `json:"Objects"`
}

type PermissionSetData struct {
	ServiceName    string          `json:"ServiceName"`
	PermissionSets []PermissionSet `json:"PermissionSets"`
}

type TempPermission struct {
	models.Permission
}

func (TempPermission) TableName() string {
	return "TempPermission"
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

		qeury := `
		CREATE TEMPORARY TABLE TempPermission( 
			ServiceName varchar(64) DEFAULT NULL,
			Name varchar(64) DEFAULT NULL,
			Action varchar(64) DEFAULT NULL,
			Object varchar(64) DEFAULT NULL,
			UNIQUE KEY ServiceNameNameActionObject (
			ServiceName, Name, Action, Object)
		  );`

		if err = tx.Exec(qeury).Error; err != nil {
			fmt.Println(err)
			return
		}

		if err = tx.CreateInBatches(&permissions, 100).Error; err != nil {
			fmt.Println(err)
			return
		}

		if err = tx.Exec(`
			INSERT IGNORE INTO Permission(ServiceName, Name, Action, Object) 
			SELECT * FROM TempPermission;
				`).Error; err != nil {
			fmt.Println(err)
			return
		}

		if err = tx.Exec(`
			DELETE FROM p USING Permission AS p
			WHERE NOT EXISTS(
				SELECT * FROM TempPermission tp
				WHERE
				p.ServiceName = tp.ServiceName AND 
				p.Name = tp.Name AND
				p.Action = tp.Action AND
				p.Object = tp.Object
			);
				`).Error; err != nil {
			return
		}

		if err = tx.Exec(`
			DROP TABLE TempPermission;
				`).Error; err != nil {
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
		Where("ID = ?", id).
		Count(&count)
	if count == 0 {
		return permission,
			errors.New("item dosen't exist")
	}
	err = db.Model(permission).
		Updates(permission).Error

	db.Where("ID = ?", id).First(&permission)
	return permission, err
}

func (db *DBORM) DeletePermission(
	id int,
) (
	permission models.Permission,
	err error,
) {
	db.Where("ID = ?", id).First(&permission)
	return permission, db.Delete(&permission).Error
}
