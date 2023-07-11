package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var motor Customer
	motor.Name = "fauzan"
	motor.Address = "bekasi"
	motor.Age = 26

	fmt.Println(motor)
}
