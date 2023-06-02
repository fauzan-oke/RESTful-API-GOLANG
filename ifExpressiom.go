package main

import "fmt"

func main() {
	name := "fauzasn"

	if name == "fauzan" {
		fmt.Println(name)
	} else if name == "hani" {
		fmt.Println(name)
	} else {
		fmt.Println("hmm im not sure")
	}

	if length := len(name); length > 10 {
		fmt.Println("nama kepanjangan")
	} else {
		fmt.Println("benar")
	}
}
