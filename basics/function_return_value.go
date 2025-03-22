package main

import "fmt"

func sayHello( name string) string {
	// buat variable hello 
	hello:= "Halo: " + name
	// wajib pakai kata kunci return untuk mengembalikan data nya
	return hello
}

func main() {
	result := sayHello("eko")
	fmt.Println(result)

	fmt.Println(sayHello("afan"))
}