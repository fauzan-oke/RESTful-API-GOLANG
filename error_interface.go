package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagian dengan Nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	var contohError error = errors.New("ups error!")
	fmt.Println(contohError.Error())

	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("hasil", hasil)
	} else {
		fmt.Println("error", err.Error())
	}
}
