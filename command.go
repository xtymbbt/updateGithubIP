package main

import (
	"net"
	"time"
)

func pingTest(githubIP string) bool {
	results := []bool{false, false}
	results[0] = test(githubIP, "443")
	results[1] = test(githubIP, "80")
	return results[0] && results[1]
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
