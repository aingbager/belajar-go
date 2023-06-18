package main

import "fmt"

type name struct {
	name1, name2 string
}

func main() {
	var udin name

	udin.name1 = "udin"
	udin.name2 = "surudin"

	fmt.Println("namaku: ", udin.name1)
	fmt.Println("nama panjangku: ", udin.name2)
}
