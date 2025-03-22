package main

import "fmt"

func random() any {
	return false
}

func main() {
	var result any = random()
	// var resultString string = result.(string)

	// fmt.Println(resultString)

	// var resiltInt int = result.(int)
	// fmt.Println(resiltInt)

	switch value := result.(type) {
	case string :
		fmt.Println("String", value)
	case int :
		fmt.Println("Int", value)
		default :
		fmt.Println("Tipe data tidak di ketahui")
	}

}