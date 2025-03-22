package main

import "fmt"

// A struct (short for structure) is used to create a collection of members of different data types,
// into a single variable.

type Customer struct {
	Name, Address string
	Age           int
}


// ini bukan lagi func tapi method
// karena dia sudah menempel dengan si customer
func (customer Customer) sayHello(nama string) {
	fmt.Println("Halo", nama, "My Name is", customer.Name)
}

func main() {
	// 1
	var afan Customer
	fmt.Println(afan)
	afan.Name = "Shafwan Junindra"
	afan.Address = "Indonesia"
	afan.Age = 16
	fmt.Println(afan)
	fmt.Println(afan.Name)
	fmt.Println(afan.Address)
	fmt.Println(afan.Age)

	// 2

	joko := Customer {
		Name: "Joko",
		Address: "Indonesia",
		Age: 16,
	}

	fmt.Println("2", joko)

	// 3

	budi := Customer{"Budi", "Indonesia", 17}
	fmt.Println("3", budi)

	budi.sayHello("agus")

	bernad := Customer{
		Name: "bernad",
	}

	bernad.sayHello("agus")
}