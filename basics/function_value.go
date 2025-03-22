package main

import "fmt"

func getGoodBye(name string) string {
	return `GoodBye ` + name
}

func main() {
	contoh := getGoodBye
	misal := getGoodBye

	fmt.Println(contoh("afan"))
	fmt.Println(misal("shafwan"))
}