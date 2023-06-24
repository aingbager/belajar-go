package main

import "fmt"

func main() {
	var number int = 5

	fmt.Println("before:", number)

	change(&number, 10)
	fmt.Println("after:", number)
}

func change(original *int, value int) {
	*original = value
}
