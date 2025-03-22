package main

import "fmt"


type Address struct {
	City, Provincy, Country string
}

func main() {
	var address1 = Address{"Sidoarjo", "jawa timur", "indonesia"}
	var address2 = &address1 // Pointer

	address2.City = "Bandung"

	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi bandung

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}