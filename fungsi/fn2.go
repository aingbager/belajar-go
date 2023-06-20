package main

import "fmt"

func kali(a int, b int) int {
	var hasil int = a * b
	return hasil
}
func main() {
	var a, b int
	var c int

	fmt.Print("masukan angka pertama: ")
	fmt.Scanln(&a)
	fmt.Print("masukan angka kedua: ")
	fmt.Scanln(&b)
	c = kali(a, b)

	fmt.Println("hasil perkalian:", c)
}
