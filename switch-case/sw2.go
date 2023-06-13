package main

import "fmt"

func main() {
	var angka int8

	fmt.Print("masukan angka: ")
	fmt.Scanln(&angka)

	switch angka {
	case 0, 1, 2, 3:
		fmt.Println("beginner")
	case 4, 5, 6:
		fmt.Println("normal")
	case 7, 8, 9:
		fmt.Println("hard")
	case 10:
		fmt.Println("expert")
	default:
		fmt.Println("input angka angka 1 sampai 10")
	}
}
