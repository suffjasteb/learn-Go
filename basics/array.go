package main

import "fmt"

func main() {
	// [jumlah data yang bisa di tampung di array nya]
	var names [4]string

	names[0] = "shafwan"
	names[1] = "junindra"
	names[2] = "putra"
	names[3] = "pratama"

	fmt.Println(names[0])

	// [...] => tidak ingin menentukan jumlah array di awal
	var nilais = [...]int {
		90, // 0
		80, // 1
		85, // 2
		75, // 3
 	}
 
	// ketika index ke 4 tidak di isi maka default nya 0

	fmt.Println(nilais[3])
	// len => length panjang array bukan isi array. result => 4
	fmt.Println(len(nilais))

}