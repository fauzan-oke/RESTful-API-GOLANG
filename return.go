package main

import "fmt"

func getHello(name string) string {
	return "hello" + name
}

func main() {
	result := getHello("fauzan")

	fmt.Println(result)
}
