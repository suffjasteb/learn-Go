package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	// Defer di Run Di Akhir function selesai
	defer logging()
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}