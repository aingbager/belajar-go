package main

import "fmt"

func main() {
	//belajar map
	var angka map[string]int
	angka = map[string]int{}
	angka["satu"] = 1
	angka["dua"] = 2

	fmt.Println(angka["satu"])
	fmt.Println(angka["dua"])
}
