package main

import "fmt"

func main() {
	nama := "afn"

	if nama == "afan" {
		fmt.Println("hello afan")
	} else if nama == "afn" {
		fmt.Println("hello afn")
	} else {
		fmt.Println("Hi Boleh Kenalan?")
	}

	// if short statemen
	if length := len(nama); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}