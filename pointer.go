package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeAddress(address *Address) {
	address.Country = "Amerixa"
}

func main() {
	address1 := Address{"subang", "jawa barat", "indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "bekasi"

	// address2 = &Address{"Malang", "JaTim", "Indo"}
	*address2 = Address{"Malang", "JaTim", "Indo"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	address4.City = "Cinere"
	fmt.Println(address4)

	alamat := Address{
		City:     "depok",
		Province: "Tangsel",
		Country:  "",
	}

	changeAddress(&alamat)
	fmt.Println(alamat)
}
