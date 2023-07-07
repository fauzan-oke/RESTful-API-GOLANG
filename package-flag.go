package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "put your database host")

	flag.Parse()

	fmt.Println(*host)
}
