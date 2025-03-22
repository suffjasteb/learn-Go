package main

import "fmt"

func main() {
	// stock berjumlah 1
	// stock := 1

	// perulangan jika stock kurang dari sama dengan 10, maka bernilai true dan kode di dalam {} akan terus di eksekusi
	// for stock <= 10 {
	// 	fmt.Println("Perulangan ke ", stock)
	// 	stock++
	// }

	// perulangan for dengan statemen jadi lebih singkat
	for stock := 1; stock <= 10; stock++ {
		fmt.Println("Perulangan ke ", stock)
	}
	fmt.Println("Selesai")
	// for bisa di gunakan untuk iterasi terhadap semua data di collection.
	// data collection contohnya Array, slice, map.

	// slice
	names := []string{"shafwan", "junindra", "putra"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	} 

	// ambil index (range)
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}
}