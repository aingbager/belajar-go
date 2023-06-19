package main

import "fmt"

func main() {

	var name = []string{"merah", "kuning", "hijau", "biru"}

	//ubah index ke 1 yaitu kuning menjadi ungu
	name[1] = "ungu"

	fmt.Println(name[1])
}
