package main

import "fmt"

type Filter func(string) string

// func sayHello(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("hello", nameFiltered)
// }

func sayHello(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHello("fauzan", spamFilter)
	sayHello("anjing", spamFilter)
}
