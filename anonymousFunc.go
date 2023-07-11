package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("youre blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("fauzan", blacklist)
	registerUser("anjing", blacklist)

	registerUser("hani", func(name string) bool {
		return name == "hahi"
	})
}
