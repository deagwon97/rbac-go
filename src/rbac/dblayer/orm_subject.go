package dblayer

import (
	accountModels "rbac-go/account/models"
	"rbac-go/common/paginate"
)

type SubjectsStatus struct {
	UserID    int  `json:"id"`
	IsAllowed bool `json:"is_allowed"`
}

type SubjectsStatusPage struct {
	Count        int             `json:"count"`
	NextPage     string          `json:"next"`
	PreviousPage string          `json:"previous"`
	List         []SubjectStatus `json:"results"`
}

func (db *DBORM) GetSubjectsStatusPage(
	roleID int, page int, pageSize int, hostUrl string,
) (
	subjectPage SubjectsStatusPage, err error,
) {

	var count int64
	db.Model(&accountModels.User{}).Count(&count)

	page, pageSize, nextPage, previousPage :=
		paginate.GetPageInfo(page, pageSize, hostUrl, count)
	subjectPage.Count = int(count)
	subjectPage.NextPage = nextPage
	subjectPage.PreviousPage = previousPage

	var subjectIDList []int
	err = db.
		Table("user").
		Select("id").
		Order("id desc").
		Scopes(paginate.Paginate(page, pageSize)).
		Find(&subjectIDList).
		Error

	subjectOfRole := SubjectsOfRole{
		RoleID:        roleID,
		SubjectIDList: subjectIDList,
	}
	subjectStatusOfRole, err := db.CheckSubjectIsAllowed(subjectOfRole)
	subjectPage.List = subjectStatusOfRole.List

	return subjectPage, err
}
