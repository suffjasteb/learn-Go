package main

import "fmt"

func main() {
	var nama1 string = "afan"
	var nama2 string = "afan"

	// jika nama 1 sama dengan nama 2 => true
	var resultTrue bool = nama1 == nama2
	// jika nama 1 tidak sana dengan nama 2 => false
	var resultFalse bool = nama1 != nama2

	fmt.Println(resultTrue)
	fmt.Println(resultFalse)
}