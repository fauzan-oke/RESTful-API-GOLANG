package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	hostname, err := os.Hostname()

	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("error", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	pass := os.Getenv("APP_PASS")

	fmt.Println(username)
	fmt.Println(pass)
}
