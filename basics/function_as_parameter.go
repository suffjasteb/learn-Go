package main

import "fmt"

type Filter func(string) string

// bisa pakai type kayak alias biar ga panjang di parameter
func sayHelloWithFilter(name string, filter Filter) {
	filterName := filter(name)
	fmt.Println("Hello", filterName)
}

func spamFilter(name string) string {
	if (name == "anjing") {
		return `...`
	} else {
		return name;
	}
}


func main() {
	sayHelloWithFilter("afan", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("anjing", filter);
}