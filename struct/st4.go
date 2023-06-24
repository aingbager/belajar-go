package main

import "fmt"

type angka struct {
	angka1 int8
	angka2 int8
}

func main() {

	var angka angka
	angka.angka1 = 4
	angka.angka2 = 7

	fmt.Println("angka jumlah", angka.angka1+angka.angka2)

}
