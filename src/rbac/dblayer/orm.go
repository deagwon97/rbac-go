package dblayer

import (
	"database/sql"
	"fmt"

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

func (db *DBORM) GetObjects(
	subjectID int,
	permissionName string,
	permissionAction string,
) (
	objects []string,
	err error,
) {
	var permissionObject []models.PermissionObject

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

func (db *DBORM) GetRoles(
	page int, pageSize int, hostUrl string,
) (
	role []models.Role, err error,
) {
	return role, err
}
