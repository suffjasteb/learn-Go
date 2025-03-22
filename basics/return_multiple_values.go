package main

import "fmt"

// kalo mau mengembalikan lebih dari 1 data bisa gunakan multiple
func getFullName() (string, string, int) {
	return "eko", "khanedy", 27
}


func main() {
	// // kalo mau simpan data dari func getFullName bisa simpan ke 2 variable
	// firstName, lastName, age := getFullName()
	// // print firstName lastName
	// fmt.Println(firstName, lastName, age)	

	// ignore value pakai => _ (underscore)
	firstName, _, age := getFullName()
	fmt.Println(firstName, age)

	
	// code progam slice parameter
	numbers := []int{10,10,10,10}
	// tambahkan titik tiga nanti otomatis di convert ke variable argument
	fmt.Println(sumAll(numbers...))

}