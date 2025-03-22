package main

import "fmt"

// kalo tipe datanya sama bisa langsung di pojok aja deklarasiin tipe datanya
func getCompleteName() (firstName, lastName string) {
	firstName = "Shafwan"
	lastName = "junindra"

	return firstName, lastName
}

func main() {
	// artinya yang a adalah firstname dan b adalah lastname
	a, b := getCompleteName()
	fmt.Println(a, b)
}