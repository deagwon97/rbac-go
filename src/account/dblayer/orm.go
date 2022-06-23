package dblayer

import (
	"rbac-go/account/models"
	"rbac-go/database"

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
		&models.User{},
	); err != nil {
		return nil, err
	}
	return db, err
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
