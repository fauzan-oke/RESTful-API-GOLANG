package main

import "fmt"

func getFullName() (string, string) {
	return "fauzan", "fakhri"
}

func getCompletedName() (first, scond, third string) {
	first = "1"
	scond = "2"
	third = "3"
	return
}

func main() {
	firstname, lastname := getFullName()
	fmt.Println(firstname, lastname)
	fmt.Println(firstname)
	fmt.Println(lastname)

	firstnames, _ := getFullName()
	fmt.Println(firstnames)

	first, scond, third := getCompletedName()
	fmt.Println(first, scond, third)
}
