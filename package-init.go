package main

import (
	"RESTful-API-GOLANG/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
