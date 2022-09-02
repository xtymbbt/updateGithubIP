package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var (
	FilePath       string
	TimeoutSeconds int64
	ipChan         = make(chan *string, 1024)
	successChan    = make(chan bool, 0)
	completeChan   = make(chan bool, 0)
	wg             sync.WaitGroup
)

func initConfig() {
	FilePath = "C:\\Windows\\System32\\drivers\\etc\\hosts"
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
			FilePath = strings.Replace(string(a), "filePath=", "", -1)
			FilePath = strings.Replace(FilePath, "\"", "", -1)
		}
		if strings.Contains(string(a), "timeoutSeconds") {
			TimeoutSeconds, err = strconv.ParseInt(strings.Replace(string(a), "timeoutSeconds=", "", -1), 10, 64)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println(FilePath)
}

func main() {
	initConfig()
	list, err := getIPList()
	if err != nil {
		fmt.Printf("Error occured in getIPList: %#v\n", err)
		return
	}
	//bannedIPs := readBannedIP()
	//var bannedIPMap = make(map[string]bool, 0)
	//for i := 0; i < len(bannedIPs); i++ {
	//	if _, ok := bannedIPMap[bannedIPs[i]]; !ok {
	//		bannedIPMap[bannedIPs[i]] = true
	//	}
	//}
	//kinds := []string{"hooks", "web", "api", "git", "packages", "pages", "importer", "actions", "dependabot"}
	kinds := []string{"web", "api", "git"}
	githubIPs := make(map[string]bool, 0)
	for i := 0; i < len(kinds); i++ {
		ipKind := list[kinds[i]]
		switch reflect.TypeOf(ipKind).Kind() {
		case reflect.Slice, reflect.Array:
			s := reflect.ValueOf(ipKind)
			for i := 0; i < s.Len(); i++ {
				githubIP := s.Index(i).Interface().(string)
				githubIP = strings.Split(githubIP, "/")[0]
				if strings.Contains(githubIP, ".") {
					if x := strings.Split(githubIP, ".")[3]; x == "0" {
						continue
					}
				}
				//if _, ok := bannedIPMap[githubIP]; ok {
				//	continue
				//}
				githubIPs[githubIP] = true
			}
		default:
			fmt.Println("Getting " + kinds[i] + " IP list failed.")
		}
	}
	go TestAll(githubIPs)
	go Write()
	select {
	case <-successChan:
		fmt.Println("Successfully written into hosts, you can now visit github.com, congratulations!")
	case <-completeChan:
		fmt.Println("We have tested all IPs, sadly, no ip available.")
	}
	close(successChan)
	close(completeChan)
}

func TestAll(githubIPs map[string]bool) {
	for githubIP := range githubIPs {
		wg.Add(1)
		go Test(githubIP)
	}
	wg.Wait()
	completeChan <- true
	close(ipChan)
}

func Test(githubIP string) {
	fmt.Println("Testing github ip: ", githubIP)
	if pingTest(githubIP) {
		fmt.Println("githubIP: ", githubIP, " test success.")
		ipChan <- &githubIP
	} else {
		fmt.Println("githubIP: ", githubIP, " test failed.")
	}
	wg.Done()
}

func Write() {
	for githubIPp := range ipChan {
		githubIP := *githubIPp
		err := writeHosts(githubIP)
		if err != nil {
			fmt.Printf("Error occured in write Hosts: %#v\n", err)
			continue
		}
		successChan <- true
		close(ipChan)
	}
}
