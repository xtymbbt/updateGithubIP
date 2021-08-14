package main

import (
	"net"
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
