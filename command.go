package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func pingTest(githubIP string) bool {
	// 测试三次
	for i := 0; i < 3; i++ {
		if !test(githubIP, "443") {
			return false
		}
		if !test(githubIP, "80") {
			return false
		}
	}
	if !testHttp(githubIP) {
		return false
	}
	return true
}

func testHttp(githubIP string) bool {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://" + githubIP + ":443")
	if err != nil {
		fmt.Printf("error occured in testing http: %#v\nErr: %#v\n", err, err.Error())
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error occured in read http result: %#v\n", err)
		return false
	}
	startIdx := 0
	endIdx := 0
	for i := 0; i < len(body); i++ {
		if body[i] == '<' {
			startIdx = i
		} else if body[i] == '>' {
			endIdx = i
			break
		}
	}
	content := string(body[startIdx : endIdx+1])
	print(content)
	if content != "<!DOCTYPE html>" {
		return false
	}
	return true
}
func test(IP string, port string) bool {
	address := net.JoinHostPort(IP, port)
	// 1 秒超时
	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		return false
	} else {
		if conn != nil {
			_ = conn.Close()
			return true
		} else {
			return false
		}
	}
}
