package main

import "fmt"

func opps(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}

}
func main() {
	var data interface{} = opps(3)
	fmt.Println(data)
}
