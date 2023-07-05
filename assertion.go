package main

import "fmt"

func random() interface{} {
	return 10
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("value", value, "is string")
	case int:
		fmt.Println("value", value, "is int")
	default:
		fmt.Println("unknown")
	}

	// resultInt := result.(int)
	// fmt.Println(resultInt)

}
