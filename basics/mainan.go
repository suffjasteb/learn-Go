package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func profil(firstName string, lastName string, age int)  (string, string, int) {
	firstName = "shafwan"
	lastName = "junindra"
	age = 16
	return firstName, lastName, age;
}

// nill atau Null
func NewMap(name string) map[string] string {
	if name == "" {
		return nil
	} else {
		return map[string] string {
			"name": name,
		}
	}
}

type student struct {
	name string
	grade int
}


func main() {
	firstName, lastName, age := profil("shafwan", "junindra", 16)
	result := sum(5, 8)
	fmt.Println(`Hasil 5 + 8 adalah = `, result)
	fmt.Println("nama saya : ",firstName, lastName,"umur saya : ", age)

	count := 10

	for i := 0; count >= i; i++ {
		fmt.Println("perulangan ke: ", i)
	}

	fmt.Println("---------")

	// break
	for i := 0;count >= i ; i++ {
		if (i > 3) {
			break
		}
		fmt.Println( "perulangan ke: ", i)
	}

	fmt.Println("--------")
	// continue
	for i := 0;count >= i ; i++ {
		if (i % 2 == 1 ) {
			continue
		}
		fmt.Println( "perulangan ke: ", i)
	}

	data := NewMap("mainan nil")
	fmt.Println(data["name"])

	// 

	var s1 student

	s1.name = "Shafwan Junindra"
	s1.grade = 2

	fmt.Println("Name: ", s1.name )
	fmt.Println("Grade: ", s1.grade )

}