package dblayer

import (
	"rbac-go/database"
	"rbac-go/rbac/models"

	"gorm.io/gorm"
)

type DBORM struct {
	*gorm.DB
}

// DBORM의 생성자
func NewORM() (db *DBORM, err error) {
	gormDB, err := database.CreateGormDB()
	db = &DBORM{DB: gormDB}

	if err := db.AutoMigrate(
		&models.SubjectAssignment{},
		&models.Role{},
		&models.PermissionAssignment{},
		&models.Permission{},
	); err != nil {
		return nil, err
	}
	return db, err
}
