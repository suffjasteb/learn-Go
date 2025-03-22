package main

import "fmt"

func sayHelloTo(name string, age int) {
	fmt.Println("Hello", name, "umur mu :", age, "ya?")
}

func main() {
	sayHelloTo("Shafwan junindra", 16)
	sayHelloTo("Shafwan putra", 26)
}