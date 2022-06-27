package dblayer

import (
	accountModels "rbac-go/account/models"
	"rbac-go/common/paginate"
)

type SubjectsStatus struct {
	UserID    int  `json:"ID"`
	IsAllowed bool `json:"IsAllowed"`
}

type SubjectsStatusPage struct {
	Count        int             `json:"Count"`
	NextPage     string          `json:"NextPage"`
	PreviousPage string          `json:"PreviousPage"`
	Results      []SubjectStatus `json:"Results"`
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
		Table("User").
		Select("ID").
		Order("ID desc").
		Scopes(paginate.Paginate(page, pageSize)).
		Find(&subjectIDList).
		Error
	if err != nil {
		return
	}

	subjectOfRole := SubjectsOfRole{
		RoleID:        roleID,
		SubjectIDList: subjectIDList,
	}
	subjectStatusOfRole, err := db.CheckSubjectIsAllowed(subjectOfRole)
	subjectPage.Results = subjectStatusOfRole.List

	return subjectPage, err
}
