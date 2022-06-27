package test

import (
	"fmt"
	"strconv"
	"testing"

	"rbac-go/common/test"
	"rbac-go/rbac/dblayer"
	"rbac-go/rbac/rest"
)

func parseId(data interface{}) (IDstring string) {
	roleMap1, _ := data.(map[string]interface{})
	IDfloat64 := roleMap1["ID"]
	IDint := int(IDfloat64.(float64))
	IDstring = strconv.Itoa(IDint)
	return IDstring
}

func Crud(
	hostUrl string,
	itemName string,
	data1 interface{},
	data2 interface{},
	data3 interface{},
	changeData interface{},
) {
	fmt.Println("---- 생성 테스트-----")
	var resData interface{}
	resData = test.Post(hostUrl+"/rbac/"+itemName, data1)
	dataIDStr1 := parseId(resData)
	resData = test.Post(hostUrl+"/rbac/"+itemName, data2)
	dataIDStr2 := parseId(resData)
	resData = test.Post(hostUrl+"/rbac/"+itemName, data3)
	dataIDStr3 := parseId(resData)

	fmt.Println("---- 목록 테스트-----")
	test.Get(hostUrl + "/rbac/" + itemName + "/list")

	fmt.Println("---- 정보 수정 테스트-----")
	test.Patch(hostUrl+"/rbac/"+itemName+"/"+"-1", changeData)
	test.Patch(hostUrl+"/rbac/"+itemName+"/"+dataIDStr3, changeData)

	fmt.Println("---- 목록 조회 -----")
	test.Get(hostUrl + "/rbac/" + itemName + "/list")

	fmt.Println("---- 삭제 테스트-----")
	test.Delete(hostUrl + "/rbac/" + itemName + "/" + dataIDStr1)
	test.Delete(hostUrl + "/rbac/" + itemName + "/" + dataIDStr2)
	test.Delete(hostUrl + "/rbac/" + itemName + "/" + dataIDStr3)
}

func TestAddRoles(t *testing.T) {
	hostUrl := "https://rbac.dev.deagwon.com"
	itemName := "role"

	data1 := dblayer.RoleData{
		Name:        "관리자1",
		Description: "관리자 그룹",
	}
	data2 := dblayer.RoleData{
		Name:        "중간관리자2",
		Description: "중간관리자 그룹",
	}
	data3 := dblayer.RoleData{
		Name:        "일반사용3",
		Description: "일반사용자 그룹",
	}
	data4 := dblayer.RoleData{
		Name:        "일반사용자14",
		Description: "일반사용자 그룹",
	}
	data5 := dblayer.RoleData{
		Name:        "일반사용자25",
		Description: "일반사용자 그룹",
	}
	data6 := dblayer.RoleData{
		Name:        "일반사용자35",
		Description: "일반사용자 그룹",
	}

	test.Post(hostUrl+"/rbac/"+itemName, data1)

	test.Post(hostUrl+"/rbac/"+itemName, data2)

	test.Post(hostUrl+"/rbac/"+itemName, data3)

	test.Post(hostUrl+"/rbac/"+itemName, data4)

	test.Post(hostUrl+"/rbac/"+itemName, data5)

	test.Post(hostUrl+"/rbac/"+itemName, data6)

}

func TestRoleCrud(t *testing.T) {
	hostUrl := "https://rbac.dev.deagwon.com"
	itemName := "role"

	data1 := dblayer.RoleData{
		Name:        "관리자",
		Description: "관리자 그룹",
	}
	data2 := dblayer.RoleData{
		Name:        "중간관리자",
		Description: "중간관리자 그룹",
	}
	data3 := dblayer.RoleData{
		Name:        "일반사용자",
		Description: "일반사용자 그룹",
	}

	changeData := dblayer.RoleData{
		Name:        "일반사용자2",
		Description: "일반사용자 그룹2",
	}

	Crud(hostUrl, itemName,
		data1, data2, data3,
		changeData)
}

func TestPermissionCrud(t *testing.T) {
	hostUrl := "https://rbac.dev.deagwon.com"
	itemName := "permission"

	data1 := dblayer.PermissionData{
		ServiceName: "로그인서비스",
		Name:        "관리자",
		Action:      "조회",
		Object:      "일반 사용자2",
	}
	data2 := dblayer.PermissionData{
		ServiceName: "로그인서비스",
		Name:        "관리자",
		Action:      "조회",
		Object:      "일반 사용자3",
	}
	data3 := dblayer.PermissionData{
		ServiceName: "로그인서비스",
		Name:        "관리자",
		Action:      "조회",
		Object:      "일반 사용자4",
	}
	changeData := dblayer.PermissionData{
		ServiceName: "로그인서비스",
		Name:        "관리2자",
		Action:      "조회2",
		Object:      "일반 사용자4",
	}

	Crud(hostUrl, itemName,
		data1, data2, data3,
		changeData)

}

func TestGetPermissionList(t *testing.T) {
	hostUrl := "https://rbac.dev.deagwon.com"
	itemName := "permission"
	test.Get(hostUrl + "/rbac/" + itemName + "/list")
}

func TestGetRoleList(t *testing.T) {
	hostUrl := "https://rbac.dev.deagwon.com"
	itemName := "role"
	test.Get(hostUrl + "/rbac/" + itemName + "/list")
}

func TestDeleteData(t *testing.T) {

	db, _ := dblayer.NewORM()

	result := map[string]interface{}{}

	db.Raw("DELETE FROM PermissionAssignment").Scan(&result)
	fmt.Println(result)
	db.Raw("DELETE FROM SubjectAssignment").Scan(&result)
	fmt.Println(result)
	db.Raw("DELETE FROM Role").Scan(&result)
	fmt.Println(result)
	db.Raw("DELETE FROM Permission").Scan(&result)
	fmt.Println(result)

}

func TestAddRole(t *testing.T) {
	hostUrl := "https://rbac.dev.deagwon.com"

	// Role 생성
	var roleRes interface{}
	roleData := dblayer.RoleData{
		Name:        "관리자",
		Description: "관리자 그룹",
	}
	roleRes = test.Post(hostUrl+"/rbac/role", roleData)
	roleID := parseId(roleRes)
	roleIDInt, _ := strconv.Atoi(roleID)

	// Permission 생성
	var permissionRes interface{}
	permissionData := dblayer.PermissionData{
		ServiceName: "bdg-blog",
		Name:        "관리2자",
		Action:      "조회asdf2",
		Object:      "asdf일반 사용자4",
	}
	permissionRes = test.Post(hostUrl+"/rbac/permission", permissionData)
	permissionID := parseId(permissionRes)
	permissionIDInt, _ := strconv.Atoi(permissionID)

	// Subject Assignment
	var saRes interface{}
	saData := dblayer.SubjectAssignmentData{
		SubjectID: 0,
		RoleID:    roleIDInt,
	}
	saRes = test.Post(hostUrl+"/rbac/subject-assignment", saData)
	saID := parseId(saRes)
	print(saID)

	// Permission Assignment
	var paRes interface{}
	paData := dblayer.PermissionAssignmentData{
		PermissionID: permissionIDInt,
		RoleID:       roleIDInt,
	}
	paRes = test.Post(hostUrl+"/rbac/permission-assignment", paData)
	paID := parseId(paRes)
	print(paID)

	permissionQeury := rest.PermissionQuery{
		SubjectID:   0,
		ServiceName: "bdg-blog",
		Name:        "관리2자",
		Action:      "조회asdf2",
	}

	test.Post(hostUrl+"/rbac/permission"+"/objects", permissionQeury)

	// Permission 삭제
	test.Delete(hostUrl + "/rbac/permission-assignment/" + paID)

	// Permission 삭제
	test.Delete(hostUrl + "/rbac/subject-assignment/" + saID)

	// Permission 삭제
	test.Delete(hostUrl + "/rbac/permission/" + permissionID)

	// Role 삭제
	test.Delete(hostUrl + "/rbac/role/" + roleID)
}

func TestGetSubjectsOfRolePage(t *testing.T) {
	hostUrl := "https://rbac.dev.deagwon.com"
	roleID := 339
	fullUrl := fmt.Sprintf(
		"%s/rbac/role/%d/subject", hostUrl, roleID)
	test.Get(fullUrl)
}

func TestPermissionsOfRolePage(t *testing.T) {
	hostUrl := "https://rbac.dev.deagwon.com"
	roleID := 339
	fullUrl := fmt.Sprintf(
		"%s/rbac/role/%d/permission", hostUrl, roleID)
	test.Get(fullUrl)
}

func TestAddPermissionSet(t *testing.T) {
	hostUrl := "https://rbac.dev.deagwon.com"

	// Permission 생성
	permissionSet1 := dblayer.PermissionSet{
		Name:    "게시판",
		Actions: []string{"상세조회", "목록조회", "수정", "삭제"},
		Objects: []string{"공지", "자유", "비밀"},
	}

	permissionSet2 := dblayer.PermissionSet{
		Name:    "채팅",
		Actions: []string{"상세조회", "목록조회", "수정", "삭제"},
		Objects: []string{"VIP", "도매"},
	}

	permissionSetData := dblayer.PermissionSetData{
		ServiceName: "bdg블로그",
		PermissionSets: []dblayer.PermissionSet{
			permissionSet1,
			permissionSet2,
		},
	}

	res := test.Post(hostUrl+"/rbac/permission/set", permissionSetData)
	fmt.Println(res)

	db, _ := dblayer.NewORM()

	result := map[string]interface{}{}
	db.Raw("DELETE FROM permission").Scan(&result)
	fmt.Println(result)
}

func TestGetObjectList(t *testing.T) {
	hostUrl := "https://rbac.dev.deagwon.com"
	itemName := "permission"

	data := rest.PermissionQuery{
		SubjectID:   3,
		ServiceName: "aaa",
		Name:        "asdf",
		Action:      "as",
	}

	test.Post(hostUrl+"/rbac/"+itemName+"/objects", data)
}
