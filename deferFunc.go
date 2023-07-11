package main

import "fmt"

func logging() {
	fmt.Println("memanggil func logging")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("run app")
	result := 10 / value
	fmt.Println("result", result)
}

func main() {
	runApplication(10)
	// runApplication(0)

}
