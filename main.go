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
		fmt.Printf("Error occured: %#v\n", err)
		return
	}
	bannedIPs := readBannedIP()
	var bannedIPMap = make(map[string]bool, 0)
	for i := 0; i < len(bannedIPs); i++ {
		if _, ok := bannedIPMap[bannedIPs[i]]; !ok {
			bannedIPMap[bannedIPs[i]] = true
		}
	}
	git := list["git"]
	switch reflect.TypeOf(git).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(git)
		for i := 0; i < s.Len(); i++ {
			githubIP := s.Index(i).Interface().(string)
			githubIP = strings.Split(githubIP, "/")[0]
			if x := strings.Split(githubIP, ".")[3]; x == "0" {
				continue
			}
			if _, ok := bannedIPMap[githubIP]; ok {
				continue
			}
			fmt.Println("Testing github ip: ", githubIP)
			test := pingTest(githubIP)
			if test {
				err := writeHosts(githubIP)
				if err != nil {
					fmt.Printf("Error occured: %#v\n", err)
					return
				} else {
					err = writeBannedIP(githubIP)
					if err != nil {
						fmt.Printf("Error occured: %#v\n", err)
						return
					} else {
						fmt.Printf("Write into hosts succeeded.\nPress any key to exit...\n")
						b := make([]byte, 1)
						os.Stdin.Read(b)
						return
					}
				}
			} else {
				err = writeBannedIP(githubIP)
				if err != nil {
					fmt.Printf("Error occured: %#v\n", err)
					return
				} else {
					fmt.Println("Test failed.")
					fmt.Println("Beginning another test:")
				}
			}
		}
		fmt.Println("We have tested all IPs.")
	default:
		fmt.Println("Getting IP list failed.")
	}
}
