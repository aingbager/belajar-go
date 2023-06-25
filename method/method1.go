package main

import "fmt"

type pelajar struct {
	name string
	usia int
}

func (s pelajar) sayHello() {
	fmt.Println("hello", s.name)
}

func (s pelajar) ubahName(name string) string {
	name = s.name
	return name
}
func main() {

	var udin = pelajar{"udin", 24}
	udin.sayHello()
	fmt.Println("name:", udin.name)
	fmt.Println("usia:", udin.usia)

	var otong = udin.ubahName("otong")

	//ubahname
	udin.ubahName("otong")

}
