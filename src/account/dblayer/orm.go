package dblayer

import (
	"database/sql"

	"rbac-go/account/models"

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

func (db *DBORM) AddUser(
	user models.User,
) (
	models.User, error,
) {
	err := db.Create(&user).Error
	return user, err
}

func (db *DBORM) GetPassword(
	loginID string) (models.User, error) {
	var user models.User
	err := db.Model(&models.User{}).
		Select("id", "login_id", "password").
		Where("login_id = ?", loginID).Scan(&user).
		Error
	return user, err
}

type UserIDName struct {
	ID   int    `gorm:"id"    json:"id"`
	Name string `gorm:"name"  json:"name"`
}

func (db *DBORM) GetUserListName(userIDList []int) (userIDName []UserIDName, err error) {

	err = db.
		Table("user").
		Raw(`
		SELECT id, name FROM user where id in ? ORDER BY id DESC;
		`, userIDList).
		Scan(&userIDName).
		Error

	return
}
