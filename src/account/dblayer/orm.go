package dblayer

import (
	"database/sql"
	"strconv"

	"rbac/account/models"

	"time"

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

func (db *DBORM) GetAllContents(
	page int, pageSize int, hostUrl string) (
	contents models.ContentList, err error) {
	var count int64
	db.Model(&models.ContentItem{}).Count(&count)

	page, pageSize, nextPage, previousPage := GetPageInfo(page, pageSize, hostUrl, count)
	contents.Count = int(count)
	contents.NextPage = nextPage
	contents.PreviousPage = previousPage

	err = db.
		Select("content_id", "title", "summary").
		Order("content_id desc").
		Scopes(Paginate(page, pageSize)).
		Find(&contents.Results).
		Error

	return contents, err
}

func (db *DBORM) GetContent(id int) (content models.Content, err error) {
	return content, db.First(&content, id).Error
}

func (db *DBORM) AddContent(contentData models.ContentData) (content models.Content, err error) {
	content.User = 2
	loc, _ := time.LoadLocation("Asia/Seoul")
	kst := time.Now().In(loc)
	content.CreatedAt = kst.String()

	content.Title = contentData.Title
	content.Summary = contentData.Summary
	content.Content = contentData.Content
	err = db.Create(&content).Error
	return content, err
}

func (db *DBORM) UpdateContent(
	id int, contentData models.ContentData) (
	models.Content, error) {

	contentData.ID = id

	loc, _ := time.LoadLocation("Asia/Seoul")
	kst := time.Now().In(loc)
	contentData.UpdatedAt = kst.String()
	err := db.Model(&contentData).Updates(contentData).Error

	var content models.Content
	db.Where("content_id = ?", id).First(&content)
	return content, err
}

func (db *DBORM) DeleteContent(id int) (models.Content, error) {
	var content models.Content
	db.Where("content_id = ?", id).First(&content)
	return content, db.Delete(&content).Error
}
