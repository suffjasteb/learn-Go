package main

import "fmt"

func main() {
	// alias untuk tipe data string
	type NoKtp string

	var ktpBeril NoKtp = "11199182928192"

	fmt.Println(ktpBeril)
}