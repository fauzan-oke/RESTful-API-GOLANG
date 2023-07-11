package main

import "fmt"

func getGoodBye(name string) string {
	return "good by " + name
}

func main() {
	sayBye := getGoodBye

	fmt.Println(sayBye("fauzan"))
}
