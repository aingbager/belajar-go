package main

import "fmt"

func main() {
	//belajar pointer
	var numberA int = 5
	var numberB *int = &numberA

	fmt.Println("number a =", numberA)
	fmt.Println("number a adress =", &numberA)

	fmt.Println("number b =", *numberB)
	fmt.Println("number b adress =", numberB)
}
