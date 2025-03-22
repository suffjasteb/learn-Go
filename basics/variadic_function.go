package main

import "fmt"

//Fungsi sumAll() menerima parameter yang disebut numbers, yang memiliki tipe variadic ...int.
//Artinya, bisa mengirimkan sejumlah argumen bertipe int ke fungsi ini, tanpa batasan jumlahnya
func sumAll(numbers ...int) int {
	total := 0
	//range digunakan untuk iterasi melalui setiap elemen dalam slice numbers, meskipun 
	//kita tidak menggunakan indeksnya (karena kita hanya butuh nilai dari setiap elemen).
	for _, number := range numbers {
		// Setiap elemen dalam numbers akan disimpan dalam variabel number, 
		//dan kita menambahkan nilai number ke dalam total.
		total += number
	}
	// Setelah loop selesai, fungsi akan mengembalikan nilai total, yaitu 
	//hasil penjumlahan dari semua angka yang diterima.
	return total
}

func main() {
	fmt.Println(sumAll(10, 10, 10, 10, 10))
}