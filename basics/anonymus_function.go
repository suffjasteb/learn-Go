package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are Blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("eko", blacklist)
	registerUser("anjing", blacklist)

	registerUser("beni", func(name string) bool {
		return name == "anjing"
	})
}