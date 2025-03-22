package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string {}
	// person["nama"] = "budi"
	// person["address"] = "Sidoarjo"
	// tipe nya map - tipe data nya string
	person := map[string]string {
		"nama": "Afan",
		"address": "D`Gardenia City Blok i43",
	}

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["address"])

	book := map[string]string {
		"name": "Tutorial Golang", // key
		"author": "Eko Kurniawan", //key
		"salah": "Ups", //key 
	}

	fmt.Println(book) // book => map
	// delete map - key
	delete(book, "salah")
	// ubah data di map dengan key
	book["author"] = "budi Nugraha"
	fmt.Println(book)
}