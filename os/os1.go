package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("selamat datang")
	os.Exit(1) //exit paksa
	fmt.Println("hallo")
}
