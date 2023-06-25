package main

import (
	"fmt"
	"strconv"
)

func main() {
	//belajar error

	var input string
	fmt.Print("masukan angka:")
	fmt.Scanln(&input)

	var angka int
	var err error

	angka, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(angka, "ini angka")
	} else {
		fmt.Println(input, "bukan angka")
	}
}
