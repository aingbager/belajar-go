package main

import "fmt"

type name struct {
	name string
	age  int
}

func main() {
	//deklary setruk emnggunakan type
	var udin name
	udin.name = "udin"
	udin.age = 33

	fmt.Println("my name: ", udin.name)
	fmt.Println("my age: ", udin.age)
}
