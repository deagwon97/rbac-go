package dblayer

import (
	"database/sql"
	"fmt"
	"strconv"

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
	// gorm.Open은 *gorm.DB 타입을 초기화한다.
	gormDB, err := gorm.Open(
		mysql.New(mysql.Config{Conn: sqlDB}),
		&gorm.Config{},
	)
	return &DBORM{
		DB: gormDB,
	}, err
}

type PermissionObject struct {
	PermissionObject string `gorm:"column:permission_object" json:"permission_object"`
}

func (db *DBORM) GetObjects(
	subjectID int,
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
	INNER JOIN role r
		ON r.id = pa.role_id
	INNER JOIN subject_assignment s
		ON r.id = s.role_id
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

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetPageInfo(
	page int, pageSize int, hostUrl string, count int64) (
	int, int, string, string) {

	if page == 0 {
		page = 1
	}
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	var currentCount int64
	currentCount = int64((page) * pageSize)
	var nextPage string
	if count <= currentCount {
		nextPage = ""
	} else {
		nextPage = hostUrl +
			"?page=" +
			strconv.Itoa(page+1) +
			"pageSize=" +
			strconv.Itoa(pageSize)
	}
	var previousPage string
	if page == 1 {
		previousPage = ""
	} else {
		previousPage = hostUrl +
			"?page=" +
			strconv.Itoa(page-1) +
			"pageSize=" +
			strconv.Itoa(pageSize)
	}
	return page, pageSize, nextPage, previousPage
}

type RoleList struct {
	Roles []models.Role `json:"results"`
}

func (db *DBORM) GetRoleList() (
	roleList RoleList, err error,
) {
	err = db.
		Find(&roleList.Roles).
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

	page, pageSize, nextPage, previousPage := GetPageInfo(page, pageSize, hostUrl, count)
	rolePage.Count = int(count)
	rolePage.NextPage = nextPage
	rolePage.PreviousPage = previousPage

	err = db.
		Select("id", "name", "description").
		Order("id desc").
		Scopes(Paginate(page, pageSize)).
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
	err = db.Model(&role).Updates(role).Error

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
