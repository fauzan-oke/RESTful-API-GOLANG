package main

import "fmt"

func main() {
	var month = [...]string{
		"jan",
		"feb",
		"maret",
		"apr",
		"mei",
		"juni",
		"juli",
		"august",
		"sept",
		"okt",
		"nov",
		"des",
	}

	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	newSlice := make([]string, 3, 5)
	newSlice[0] = "fauzan"
	newSlice[1] = "fakhri"
	newSlice[2] = "ri"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
}
