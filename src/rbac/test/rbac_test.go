package test

import (
	"fmt"
	"strconv"
	"testing"

	"rbac-go/common/test"
	"rbac-go/rbac/dblayer"
)

func parseId(data interface{}) (IDstring string) {
	roleMap1, _ := data.(map[string]interface{})
	IDfloat64 := roleMap1["id"]
	IDint := int(IDfloat64.(float64))
	IDstring = strconv.Itoa(IDint)
	return IDstring
}
func TestRbac(t *testing.T) {
	HostUrl := "https://rbac.dev.deagwon.com"

	fmt.Println("----role 생성 테스트-----")
	roleData1 := dblayer.RoleData{
		Name:        "관리자",
		Description: "관리자 그룹",
	}
	roleData2 := dblayer.RoleData{
		Name:        "중간관리자",
		Description: "중간관리자 그룹",
	}
	roleData3 := dblayer.RoleData{
		Name:        "일반사용자",
		Description: "일반사용자 그룹",
	}
	var data interface{}
	data = test.Post(HostUrl+"/rbac/role", roleData1)
	roleIDStr1 := parseId(data)
	data = test.Post(HostUrl+"/rbac/role", roleData2)
	roleIDStr2 := parseId(data)
	data = test.Post(HostUrl+"/rbac/role", roleData3)
	roleIDStr3 := parseId(data)

	fmt.Println("----role 목록 테스트-----")
	test.Get(HostUrl + "/rbac/role/list")

	fmt.Println("----role 없는 정보 수정 테스트-----")
	changeData := dblayer.RoleData{
		Name:        "일반사용자2",
		Description: "일반사용자 그룹2",
	}
	test.Patch(HostUrl+"/rbac/role/12", changeData)

	fmt.Println("----role 존재하는 정보 수정 테스트-----")
	changeData = dblayer.RoleData{
		Name:        "일반사용자2",
		Description: "일반사용자 그룹2",
	}
	test.Patch(HostUrl+"/rbac/role/"+roleIDStr1, changeData)
	fmt.Println("----role 목록 조회 -----")
	test.Get(HostUrl + "/rbac/role/list")

	fmt.Println("----role 삭제 테스트-----")
	test.Delete(HostUrl + "/rbac/role/" + roleIDStr1)
	test.Delete(HostUrl + "/rbac/role/" + roleIDStr2)
	test.Delete(HostUrl + "/rbac/role/" + roleIDStr3)
}
