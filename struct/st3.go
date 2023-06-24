package main

import "fmt"

type name struct {
	FrontName string
	backName  string
}

func main() {
	//belajar struct

	var udin name
	udin.FrontName = "udin"
	udin.backName = "surudin"

	fmt.Println("namaku:", udin.FrontName)
	fmt.Println("nama panjangku:", udin.backName)

}
