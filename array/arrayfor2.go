package main

import "fmt"

func main() {
	//for range in aray

	var name = [5]string{"udin", "tong", "ucup", "dodo", "messi"}

	for _, names := range name {
		fmt.Printf("elemen : %s\n", names)
	}
}
