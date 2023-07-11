package main

import "fmt"

func main() {
	counter := 0
	name := "fauzan"

	increment := func() {
		name = "hani"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(name)
}
