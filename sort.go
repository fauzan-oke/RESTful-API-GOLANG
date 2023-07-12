package main

import (
	"fmt"
	"sort"
)

type User struct {
	Nama string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {

	users := []User{
		{"fauzan", 27},
		{"Hani", 29},
		{"Hao", 35},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)

}
