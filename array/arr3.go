package main

import "fmt"

func main() {
	//belajar array

	var angka = [5]int{1, 2, 3, 4, 5}

	fmt.Println(angka[0])
	fmt.Println(angka[1])
	fmt.Println(angka[2])
	fmt.Println(angka[3])
	fmt.Println(angka[4])

	fmt.Println("panjang slot angka", len(angka))
}
