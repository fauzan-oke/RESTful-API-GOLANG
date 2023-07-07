package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("fauzan fakhri", "fauzan"))
	fmt.Println(strings.Split("fauzan fakhri", " "))
	fmt.Println(strings.ToLower("FAUZAN FAKHRI"))
	fmt.Println(strings.ToUpper("FAUZAN FAKHRI"))
	fmt.Println(strings.Trim(" fauzan fakhri ", " "))
	fmt.Println(strings.ReplaceAll("fauzan fakhri", "fauzan", "joe"))

	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	toString, err := strconv.Atoi("2000")
	if err == nil {
		fmt.Println(toString)
	} else {
		fmt.Println(err.Error())
	}

	toInt := strconv.Itoa(2000000)
	fmt.Println(toInt)

}
