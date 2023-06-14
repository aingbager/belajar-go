package main

import "fmt"

func main() {
	var angka = [...]int8{1, 23, 4, 5, 6, 7, 88}

	for i := 0; i < len(angka); i++ {
		fmt.Println("i ke", angka[i])
	}
}
