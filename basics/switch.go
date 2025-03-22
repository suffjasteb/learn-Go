package main

import "fmt"


func main() {
	nama := "afan"

	// jika hanya untuk 1 variable bisa gunakan switch
	switch nama {
	case "afan":
	fmt.Println("Hello afan")
	case "eko":
	fmt.Println("Hello eko")
	case "budi":
	fmt.Println("Hello budi")
	// default jika semua kondisi di atas tidak terpenuhi alias else
	default :
	fmt.Println("Hai Boleh kenalan?") 
	}

	// switch - bikin variable length = len(nama); length isinya panjang variable nama
	switch length := len(nama); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang");
	case false:
		fmt.Println("nama sudah benar")
	} 
}
