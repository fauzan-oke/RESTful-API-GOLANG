package main

import "fmt"

func main() {

	var names [3]string
	names[0] = "fauzan"
	names[1] = "hani"
	names[2] = "azizah"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var value = [3]int{
		7, 15, 22,
	}

	fmt.Println(value)

}
