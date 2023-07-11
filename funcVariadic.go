package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumAll(10, 10, 10, 120)
	fmt.Println(total)

	slice := []int{11, 20, 30, 40, 50}
	total = sumAll(slice...)
	fmt.Println(total)
}
