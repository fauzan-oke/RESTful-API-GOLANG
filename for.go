package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke ", counter)
		counter++
	}

	for counters := 1; counters <= 10; counters++ {
		fmt.Println("perulangan ke ", counters)
	}

	slice := []string{"fauzan", "fakhri", "hani"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	for _, value := range slice {
		// fmt.Println("index", i, "=", value)
		fmt.Println(value)
	}
}
