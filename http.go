package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getIPList() (map[string]interface{}, error) {
	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	response, err := client.Get("https://api.github.com/meta")
	if err != nil {
		fmt.Printf("err is: %v\n", err)
		return nil, err
	}
	responseMap, err := ParseResponse(response)
	if err != nil {
		return nil, err
	}
	return responseMap, nil
}

func ParseResponse(response *http.Response) (map[string]interface{}, error){
	var result map[string]interface{}
	body,err := ioutil.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}
	return result,err
}
