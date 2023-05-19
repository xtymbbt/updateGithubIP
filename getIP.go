// Package main
// File: getIP.go
// Author: Bridge Wang
// Date: 2023-05-19 16:33:51
package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func get521XueWeiHanGitHubIP() (ipHosts [][]string, err error) {
	response, err := http.Get("https://raw.hellogithub.com/hosts.json")
	if err != nil {
		return
	}
	defer response.Body.Close()
	ipHosts, err = parse521XueWeiHanGitHubIP(response.Body)
	return
}

// parse521XueWeiHanGitHubIP parses the response body of https://raw.hellogithub.com/hosts.json
// and returns the IP list.
// host.json format:
// [
// ["140.82.112.26", "alive.github.com"],
// ["140.82.112.5", "api.github.com"],
// ["185.199.110.153", "assets-cdn.github.com"],
// ["185.199.108.133", "avatars.githubusercontent.com"],
// ["185.199.108.133", "avatars0.githubusercontent.com"],
// ["185.199.108.133", "avatars1.githubusercontent.com"],
// ["185.199.111.133", "avatars2.githubusercontent.com"],
// ["185.199.108.133", "avatars3.githubusercontent.com"],
// ["185.199.108.133", "avatars4.githubusercontent.com"],
// ["185.199.108.133", "avatars5.githubusercontent.com"],
// ["185.199.108.133", "camo.githubusercontent.com"],
// ["140.82.113.21", "central.github.com"],
// ["185.199.108.133", "cloud.githubusercontent.com"],
// ["140.82.114.10", "codeload.github.com"],
// ["140.82.112.21", "collector.github.com"],
// ["185.199.108.133", "desktop.githubusercontent.com"],
// ["185.199.108.133", "favicons.githubusercontent.com"],
// ["140.82.114.3", "gist.github.com"],
// ["52.217.70.220", "github-cloud.s3.amazonaws.com"],
// ["52.216.10.75", "github-com.s3.amazonaws.com"],
// ["52.217.137.1", "github-production-release-asset-2e65be.s3.amazonaws.com"],
// ["52.216.43.25", "github-production-repository-file-5c1aeb.s3.amazonaws.com"],
// ["54.231.199.177", "github-production-user-asset-6210df.s3.amazonaws.com"],
// ["192.0.66.2", "github.blog"],
// ["140.82.113.3", "github.com"],
// ["140.82.114.18", "github.community"],
// ["185.199.109.154", "github.githubassets.com"],
// ["151.101.129.194", "github.global.ssl.fastly.net"],
// ["185.199.110.153", "github.io"],
// ["185.199.108.133", "github.map.fastly.net"],
// ["185.199.110.153", "githubstatus.com"],
// ["140.82.113.26", "live.github.com"],
// ["185.199.108.133", "media.githubusercontent.com"],
// ["185.199.108.133", "objects.githubusercontent.com"],
// ["13.107.42.16", "pipelines.actions.githubusercontent.com"],
// ["185.199.108.133", "raw.githubusercontent.com"],
// ["185.199.108.133", "user-images.githubusercontent.com"],
// ["13.107.226.40", "vscode.dev"],
// ["140.82.113.22", "education.github.com"]
// ]
func parse521XueWeiHanGitHubIP(body io.ReadCloser) ([][]string, error) {
	// 首先从 body 中读取内容
	// 然后将内容解析成 json 格式
	// 最后将 json 格式的内容解析成 [][]string
	content, err := io.ReadAll(body)
	if err != nil {
		log.Printf("ERRO: io.ReadAll failed: %v", err)
		return nil, err
	}
	var ipHosts [][]string
	err = json.Unmarshal(content, &ipHosts)
	if err != nil {
		log.Printf("ERRO: json.Unmarshal failed: %v", err)
		return nil, err
	}
	return ipHosts, nil
}
