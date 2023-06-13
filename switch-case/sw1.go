package main

import "fmt"

func main() {
	//switch case

	angka := 2

	switch angka {
	case 8:
		fmt.Println("ini angka 8")
	case 3, 4, 5:
		fmt.Println("ini angka biasa")
	default:
		fmt.Println("angka bad")
	}
}
