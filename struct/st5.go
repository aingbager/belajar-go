package main

import "fmt"

type student struct {
	name string
	usia int
}

func main() {

	var udin = student{name: "udin", usia: 34}

	var otong *student = &udin

	otong.name = "otong"

	fmt.Println("name:", udin.name)
	fmt.Println("usia:", udin.usia)

}
