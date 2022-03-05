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
	IDfloat64 := roleMap1["id"]
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
	test.Patch(hostUrl+"/rbac/"+itemName+"/"+"1", changeData)
	test.Patch(hostUrl+"/rbac/"+itemName+"/"+dataIDStr3, changeData)

	fmt.Println("---- 목록 조회 -----")
	test.Get(hostUrl + "/rbac/" + itemName + "/list")

	fmt.Println("---- 삭제 테스트-----")
	test.Delete(hostUrl + "/rbac/" + itemName + "/" + dataIDStr1)
	test.Delete(hostUrl + "/rbac/" + itemName + "/" + dataIDStr2)
	test.Delete(hostUrl + "/rbac/" + itemName + "/" + dataIDStr3)
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

func TestAddRole(t *testing.T) {
	hostUrl := "https://rbac.dev.deagwon.com"

	// Role 생성
	var roleRes interface{}
	roleData := dblayer.RoleData{
		Name:        "관리자",
		Description: "관리자 그룹",
	}
	roleRes = test.Post(hostUrl+"/rbac/role/", roleData)
	roleID := parseId(roleRes)

	// Permission 생성
	var permissionRes interface{}
	permissionData := dblayer.PermissionData{
		ServiceName: "로그인",
		Name:        "관리2자",
		Action:      "조회asdf2",
		Object:      "asdf일반 사용자4",
	}
	permissionRes = test.Post(hostUrl+"/rbac/permission/", permissionData)
	permissionID := parseId(permissionRes)

	// Subject Assignment
	var saRes interface{}
	saData := dblayer.SubjectAssignmentData{
		SubjectID: 0,
		RoleID:    1,
	}
	saRes = test.Post(hostUrl+"/rbac/subject-assignment/", saData)
	saID := parseId(saRes)
	print(saID)

	// Permission Assignment
	var paRes interface{}
	paData := dblayer.PermissionAssignmentData{
		PermissionID: 0,
		RoleID:       0,
	}
	paRes = test.Post(hostUrl+"/rbac/permission-assignment/", paData)
	paID := parseId(paRes)
	print(paID)

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
