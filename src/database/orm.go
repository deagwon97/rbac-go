package database

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBORM struct {
	*gorm.DB
}

func CreateGormDB() (gormDB *gorm.DB, err error) {
	dbengine := "mysql"
	dsn := DataSource
	sqlDB, err := sql.Open(dbengine, dsn)
	if err != nil {
		return
	}
	gormDB, err = gorm.Open(
		mysql.New(mysql.Config{Conn: sqlDB}),
		&gorm.Config{},
	)

	if err != nil {
		return
	}
	return gormDB, err
}

func NewORM() (db *DBORM, err error) {
	gormDB, err := CreateGormDB()
	db = &DBORM{DB: gormDB}
	return db, err
}
