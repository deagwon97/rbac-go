package test

import (
	"bytes"
	"fmt"
	"io"

	"encoding/json"

	"io/ioutil"

	"net/http"
)

func NotNil(err error) {
	if err != nil {
		panic(err)
	}
}

func Reqeust(
	url string,
	requestMethod string,
	dataStruct interface{},
) (
	data interface{},
) {

	var buff io.Reader
	var req *http.Request
	var err error
	if dataStruct != nil {
		dataBytes, _ := json.Marshal(dataStruct)
		buff = bytes.NewBuffer(dataBytes)
	} else {
		buff = nil
	}
	req, err = http.NewRequest(requestMethod,
		url,
		buff,
	)

	NotNil(err)
	req.Header.Set("Content-Type",
		"application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	NotNil(err)

	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	// NotNil(err)
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, respBody, "", "\t")
	if err != nil {
		fmt.Println(string(respBody))
		return
	}
	// NotNil(err)
	str := prettyJSON.String()
	println(str)

	json.Unmarshal([]byte(respBody), &data)
	return data
}

func Get(url string) (
	data interface{}) {
	data = Reqeust(url, http.MethodGet, nil)
	return data
}

func Post(url string, dataStruct interface{}) (
	data interface{}) {
	data = Reqeust(url, http.MethodPost, dataStruct)
	return data
}

func Patch(url string, dataStruct interface{}) (
	data interface{}) {
	data = Reqeust(url, http.MethodPatch, dataStruct)
	return data
}

func Delete(url string) (
	data interface{}) {
	data = Reqeust(url, http.MethodDelete, nil)
	return data
}
