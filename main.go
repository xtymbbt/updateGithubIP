package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func main() {
	list, err := getIPList()
	if err != nil {
		return
	}
	git := list["git"]
	switch reflect.TypeOf(git).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(git)
		for i := 0; i < s.Len(); i++ {
			githubIP := s.Index(i).Interface().(string)
			githubIP = strings.Split(githubIP, "/")[0]
			fmt.Println("Testing github ip: ", githubIP)
			test := pingTest(githubIP)
			if test {
				err := writeHosts(githubIP)
				if err != nil {
					fmt.Printf("Error occured: %#v\n", err)
					return
				} else {
					fmt.Printf("Write into hosts succeeded.\nWaiting Program to exit...\n")
					time.Sleep(3 * time.Second)
					return
				}
			} else {
				fmt.Println("Test failed.")
				fmt.Println("Beginning another test:")
			}
		}
		fmt.Println("We have tested all IPs.")
	default:
		fmt.Println("Getting IP list failed.")
	}
}
