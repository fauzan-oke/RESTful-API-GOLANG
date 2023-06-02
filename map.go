package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "fauzan",
		"alamat": "babelan",
	}

	person["title"] = "backend"

	fmt.Println(person)
	fmt.Println(person["name"])
}
