package paginate

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParsePageUrl(c *gin.Context) (
	page int, pageSize int, hostUrl string) {

	page, _ = strconv.Atoi(c.Query("page"))
	pageSize, _ = strconv.Atoi(c.Query("pageSize"))

	hostName := c.Request.Host + c.Request.URL.Path
	protocol := c.Request.Header.Get("X-Forwarded-Proto")
	hostUrl = protocol + "://" + hostName

	return page, pageSize, hostUrl

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

	currentCount := int64((page) * pageSize)
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

// Example
// func (db *DBORM) GetRolesPage(
// 	page int, pageSize int, hostUrl string,
// ) (
// 	rolesPage RolesPage, err error,
// ) {

// 	var count int64
// 	db.Model(&models.Role{}).Count(&count)

// 	page, pageSize, nextPage, previousPage :=
// 		paginate.GetPageInfo(page, pageSize, hostUrl, count)
// 	rolesPage.Count = int(count)
// 	rolesPage.NextPage = nextPage
// 	rolesPage.PreviousPage = previousPage

// 	err = db.
// 		Select("id", "name", "description").
// 		Order("id desc").
// 		Scopes(paginate.Paginate(page, pageSize)).
// 		Find(&rolesPage.Roles).
// 		Error

// 	return rolesPage, err
// }
