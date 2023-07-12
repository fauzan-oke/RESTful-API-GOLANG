package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Date())

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc.Local())

	layout := "2006-01-02"
	// layout := time.RFC3339
	parse, _ := time.Parse(layout, "2023-10-10")
	fmt.Println(parse)

}
