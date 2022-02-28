package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"rbac-go/rbac/dblayer"
)

func Get(url string) {
	fmt.Println("Request to ", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, respBody, "", "\t")
	if err == nil {
		str := prettyJSON.String()
		println(str)
	}
}

func Post(url string, dataStruct interface{}) interface{} {

	dataBytes, _ := json.Marshal(dataStruct)
	buff := bytes.NewBuffer(dataBytes)
	resp, err := http.Post(url, "application/json", buff)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, respBody, "", "\t")
	if err == nil {
		str := prettyJSON.String()
		println(str)
	}

	var data interface{}
	json.Unmarshal([]byte(respBody), &data)
	return data
}

func Patch(url string, dataStruct interface{}) {

	dataBytes, _ := json.Marshal(dataStruct)
	buff := bytes.NewBuffer(dataBytes)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPatch, url,
		buff,
	)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, respBody, "", "\t")
	if err == nil {
		str := prettyJSON.String()
		println(str)
	}
}

func Delete(url string) {
	fmt.Println("Request to ", url)
	req, err := http.NewRequest(
		http.MethodDelete,
		url,
		nil,
	)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, respBody, "", "\t")
	if err == nil {
		str := prettyJSON.String()
		println(str)
	}
}

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
	data = Post(HostUrl+"/rbac/role", roleData1)
	roleIDStr1 := parseId(data)
	data = Post(HostUrl+"/rbac/role", roleData2)
	roleIDStr2 := parseId(data)
	data = Post(HostUrl+"/rbac/role", roleData3)
	roleIDStr3 := parseId(data)

	fmt.Println("----role 목록 테스트-----")
	Get(HostUrl + "/rbac/role/list")

	fmt.Println("----role 없는 정보 수정 테스트-----")
	changeData := dblayer.RoleData{
		Name:        "일반사용자2",
		Description: "일반사용자 그룹2",
	}
	Patch(HostUrl+"/rbac/role/12", changeData)

	fmt.Println("----role 존재하는 정보 수정 테스트-----")
	changeData = dblayer.RoleData{
		Name:        "일반사용자2",
		Description: "일반사용자 그룹2",
	}
	Patch(HostUrl+"/rbac/role/"+roleIDStr1, changeData)
	fmt.Println("----role 목록 조회 -----")
	Get(HostUrl + "/rbac/role/list")

	fmt.Println("----role 삭제 테스트-----")
	Delete(HostUrl + "/rbac/role/" + roleIDStr1)
	Delete(HostUrl + "/rbac/role/" + roleIDStr2)
	Delete(HostUrl + "/rbac/role/" + roleIDStr3)
}
