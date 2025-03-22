package main

import "fmt"

func main () {
	
for i:= 0; i < 10; i++ {
	// i SISA BAGI alias ganjil aja yg di cetak.s
	if (i % 2 == 1) {
		continue
	}
	fmt.Println("perulangan ke :", i)
}
}