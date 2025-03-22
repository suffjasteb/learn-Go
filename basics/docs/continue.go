package main

import "fmt"

func main() {
	var count int = 10
	for i := 0; i < count; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka ", i)
	}

	name := "afan"

	age := 17

	fmt.Printf("halo namaku : %s umur ku : %d tahun", name, age)
}