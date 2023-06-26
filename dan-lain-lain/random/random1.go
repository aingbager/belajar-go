package main

import (
	"fmt"
	"math/rand"
)

func main() {
	acak := rand.New(rand.NewSource(10))

	fmt.Println("acak ke 1:", acak.Int())
	fmt.Println("acak ke 2:", acak.Int())
	fmt.Println("acak ke 3:", acak.Int())
}
