package paginate

import (
	"strconv"

	"gorm.io/gorm"
)

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
// 	rolePage RolePage, err error,
// ) {

// 	var count int64
// 	db.Model(&models.Role{}).Count(&count)

// 	page, pageSize, nextPage, previousPage :=
// 		paginate.GetPageInfo(page, pageSize, hostUrl, count)
// 	rolePage.Count = int(count)
// 	rolePage.NextPage = nextPage
// 	rolePage.PreviousPage = previousPage

// 	err = db.
// 		Select("id", "name", "description").
// 		Order("id desc").
// 		Scopes(paginate.Paginate(page, pageSize)).
// 		Find(&rolePage.Roles).
// 		Error

// 	return rolePage, err
// }
