package main

import "fmt"

func main() {
	name := "fauzan"

	switch name {
	case "fauzan":
		fmt.Println("hi, fauzan")
	case "hani":
		fmt.Println("hi, hani")
	default:
		fmt.Println("siapa kamu?")

	}

	//short statement
	switch length := len(name); length > 10 {
	case true:
		fmt.Println("nama kepanjangan")
	case false:
		fmt.Println("benar")
	}
}
