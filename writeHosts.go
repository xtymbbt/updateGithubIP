package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func writeHosts(githubIP string) error {
	filePath := "C:\\Windows\\System32\\drivers\\etc\\hosts"
	conf, err := os.Open("conf.txt")
	if err != nil {
		panic(err)
	}
	confBuf := bufio.NewReader(conf)
	for {
		a, _, err := confBuf.ReadLine()
		if err == io.EOF {
			break
		}
		if strings.Contains(string(a), "filePath") {
			filePath = strings.Replace(string(a), "filePath=", "", -1)
			filePath = strings.Replace(filePath, "\"", "", -1)
		}
	}
	fmt.Println(filePath)
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)
	var result = ""
	contain := []bool{false, false}
	for {
		a, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		if strings.Contains(string(a), "www.github.com") {
			result += strings.Join([]string{githubIP, "www.github.com"}, "\t") + "\n"
			contain[0] = true
		} else if strings.Contains(string(a), "github.com") {
			result += strings.Join([]string{githubIP, "github.com"}, "\t") + "\n"
			contain[1] = true
		} else {
			result += string(a) + "\n"
		}
	}
	if !contain[0] {
		result += strings.Join([]string{githubIP, "www.github.com"}, "\t") + "\n"
	}
	if !contain[1] {
		result += strings.Join([]string{githubIP, "github.com"}, "\t") + "\n"
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	err = os.Chmod(filePath, 0666)
	if err != nil {
		panic(err)
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(fw)
	_, err = w.WriteString(result)
	if err != nil {
		panic(err)
	}
	err = w.Flush()
	if err != nil {
		panic(err)
	}
	return err
}

func readBannedIP() []string {
	filePath := "bannedIP.txt"
	//fmt.Println(filePath)
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)
	var result = make([]string, 0, 0)
	for {
		a, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		result = append(result, string(a))
	}
	err = f.Close()
	return result
}

func writeBannedIP(githubIP string) error {
	filePath := "bannedIP.txt"
	//fmt.Println(filePath)
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		result += string(a) + "\n"
	}
	result += githubIP + "\n"
	err = f.Close()
	if err != nil {
		panic(err)
	}
	//fmt.Println(result)
	err = os.Chmod(filePath, 0666)
	if err != nil {
		panic(err)
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(fw)
	_, err = w.WriteString(result)
	if err != nil {
		panic(err)
	}
	err = w.Flush()
	if err != nil {
		panic(err)
	}
	return err
}
