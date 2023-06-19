package main

import "fmt"

type name struct {
	nama_depan    string
	nama_belakang string
}

func (udin name) sayhello() {
	fmt.Println("hello", udin.nama_depan)
}

func main() {

	var udin = name{"udin", "surudin"}
	fmt.Println("my name: ", udin.nama_depan, udin.nama_belakang)
	udin.sayhello()
}
