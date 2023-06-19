package main

import "fmt"

func main() {
	//definisi slice tidak menggunkan index saat pendefinisian

	var name = []string{"merah", "kuning", "hijau", "biru"}

	fmt.Println(name[:]) //merah kuning hijau biru

	fmt.Println(name[0:2]) //merah kuning
	fmt.Println(name[1:2]) // kuning
	fmt.Println(name[2:])  //hijau biru
	fmt.Println(name[:2])  //merah kuning
	fmt.Println(name[0:0]) //kosong
}
