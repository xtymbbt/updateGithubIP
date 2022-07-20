package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	list, err := getIPList()
	if err != nil {
		fmt.Printf("Error occured in getIPList: %#v\n", err)
		return
	}
	bannedIPs := readBannedIP()
	var bannedIPMap = make(map[string]bool, 0)
	for i := 0; i < len(bannedIPs); i++ {
		if _, ok := bannedIPMap[bannedIPs[i]]; !ok {
			bannedIPMap[bannedIPs[i]] = true
		}
	}
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
				if _, ok := bannedIPMap[githubIP]; ok {
					continue
				}
				githubIPs[githubIP] = true
			}
		default:
			fmt.Println("Getting " + kinds[i] + " IP list failed.")
		}
	}
	for githubIP := range githubIPs {
		if Test(githubIP) {
			return
		}
	}
	fmt.Println("We have tested all IPs. Clear the bannedIP.txt to restart again.")
}

func Test(githubIP string) bool {
	fmt.Println("Testing github ip: ", githubIP)
	test := pingTest(githubIP)
	if test {
		err := writeHosts(githubIP)
		if err != nil {
			fmt.Printf("Error occured in write Hosts: %#v\n", err)
			return false
		} else {
			err = writeBannedIP(githubIP)
			if err != nil {
				fmt.Printf("Error occured in writeBannedIP 1: %#v\n", err)
				return false
			} else {
				fmt.Printf("Write into hosts succeeded.\nPress any key to exit...\n")
				b := make([]byte, 1)
				_, err := os.Stdin.Read(b)
				if err != nil {
					return true
				}
				return true
			}
		}
	} else {
		err := writeBannedIP(githubIP)
		if err != nil {
			fmt.Printf("Error occured in writeBannedIP 2: %#v\n", err)
			return false
		} else {
			fmt.Println("Test failed.")
			fmt.Println("Beginning another test:")
		}
	}
	return false
}
