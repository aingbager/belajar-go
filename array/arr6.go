package main

import "fmt"

func main() {
	var name [5]string

	name[0] = "udin"
	name[1] = "dudung"
	name[2] = "dodol"
	name[3] = "jojo"
	name[4] = "soso"

	fmt.Println("inde aray length:", len(name))
	fmt.Println(name[4])
}
