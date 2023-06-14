package main

import "fmt"

func main() {

	//inisaisi jumlah array tanpa menyebut sloot (...)
	var angka = [...]int8{
		1,
		2,
		3,
		4,
		5,
	}

	fmt.Println("panjang kotak", len(angka))
}
