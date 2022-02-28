package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
