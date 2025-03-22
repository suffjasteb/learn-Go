package main

import (
	"fmt"
)

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi Panic", message)
}

func runApp(error bool) {
	// ketika panic true defer teteap tereksekusi
	defer endApp()
	if error {
		panic("Ups Error")
	}
	
}

func main() {
	runApp(true)
	fmt.Println("Shafwan junindra Putra Pratama")
}