package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("f([a-z])n")

	fmt.Println(regex.MatchString("fzn"))
	fmt.Println(regex.MatchString("fauZan"))
	fmt.Println(regex.MatchString("fakhri"))

	search := regex.FindAllString("toi fzn fun fan fin yui", -1)
	fmt.Println(search)
}
