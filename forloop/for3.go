package main

import "fmt"

func main() {
	//penggunaan for tanpa argumen

	i := 0

	for {
		fmt.Println("i ke: ", i)
		i++

		if i == 10 {
			break
		}
	}
}
