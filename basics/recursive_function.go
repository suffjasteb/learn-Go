package main

import "fmt"

// pakai for lop
func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// pakai recursive => func yang memanggil diri nya sendiri
func factoriaRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factoriaRecursive(value-1)
	}
}

func main() {
	result := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1;
	fmt.Println(result)
	fmt.Println(factoriaRecursive(10))
	fmt.Println(factorialLoop(10))
}