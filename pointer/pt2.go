package main

import "fmt"

func main() {
	var numberA int = 5
	var numberB *int = &numberA

	fmt.Println("alamat memory umberA:", &numberA)
	fmt.Println("alamat memory numberB:", numberB)

	fmt.Println("isi numberA", numberA)
	fmt.Println("isi numberB", *numberB)
}
