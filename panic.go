package main

import "fmt"

func endApp() {
	fmt.Println("aplikasi selesai")
	message := recover() //untuk menaruh message error
	fmt.Println("error", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("error")
	}

	fmt.Println("aplikasi jalan")
}

func main() {
	// runApp(false)
	runApp(true)

}
