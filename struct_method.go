package main

import "fmt"

type Customers struct {
	Name, Address string
	Age           int
}

func (customer Customers) sayHello(name string) {
	fmt.Println("Hello, ", name, "my Name is", customer.Name)
}

func main() {
	var eko Customers
	eko.Name = "eko"
	eko.sayHello("fauzan")
}
