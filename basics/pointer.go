package main

import "fmt"

type Addres struct {
	City, Provincy, Country string
}

func main() {
	// addres1 := Addres{"Sidoarjo", "jawa timur", "indonesia"}
	// addres2 := addres1 // Copy

	// addres2.City = "Bandung"

	// fmt.Println(addres1.City) // tidak berubah 
	// fmt.Println(addres2.City) // berubah sidoarjo

	addres1 := Addres{"Sidoarjo", "jawa timur", "indonesia"}
	addres2 := &addres1 // Pointer

	addres2.City = "Bandung"

	fmt.Println(addres1.City) // ikut berubah menjadi sidoarjo 
	fmt.Println(addres2.City) // berubah sidoarjo
}