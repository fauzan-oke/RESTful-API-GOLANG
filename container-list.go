package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("fauzan")
	data.PushBack("fakhri")
	data.PushFront("joe")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
