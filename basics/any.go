package main

import "fmt"

func Ups() any {
	// return "Ups"
	// return 1
	return true
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}