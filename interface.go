package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHeylo(hasName HasName) {
	fmt.Println("hello", hasName.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var fauzan Person
	fauzan.Name = "fauzan"

	sayHeylo(fauzan)

	cat := Animal{
		Name: "push",
	}
	sayHeylo(cat)

}
