package main

import "fmt"

type Address struct {
	City, Provincy, Country string
}

func main() {
	// jika mau data kosong bisa pakai new
	new1 := new(Address)
	new2 := new1 // pointer

	new2.Country = "Indonesia"

	fmt.Println(new1.Country)
	fmt.Println(new2.Country)

}
