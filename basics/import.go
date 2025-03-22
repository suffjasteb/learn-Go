package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayName("budi")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa karena harus awal huruf besar untuk import

	// result2 := helper.sayGoodBye("wahyu") // tidak bisa karena harus awal huruf besar untuk import
	// fmt.Println(result2) // tidak muncul atau error
}