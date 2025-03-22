package main

import "fmt"

func main() {
	// nilai ujian saya
	var nilaiUjian = 76
	// nilai absensi saya
	var nilaiAbsensi = 10

	var lulusUjian bool = nilaiUjian > 75
	var lulusAbsensi bool = nilaiAbsensi > 80

	var lulus bool = lulusUjian && lulusAbsensi
	if (!lulus) {
		fmt.Println("Anda tidak lulus")
	} else {
		fmt.Println("Anda lulus")
	}
	fmt.Println(lulus)
}