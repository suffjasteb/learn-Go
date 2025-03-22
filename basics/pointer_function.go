package main

import "fmt"

type Address struct {
	City, Provincy, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
	address.City = "Sidoarjo"
}

func main() {
	var address = Address{}
	ChangeCountryToIndonesia(&address)

	fmt.Println(address)
}