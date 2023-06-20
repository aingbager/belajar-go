package main

import "fmt"

func main() {
	var moon map[string]int
	moon = map[string]int{}
	moon["januari"] = 50
	moon["februari"] = 20

	fmt.Println(moon["januari"])  //50
	fmt.Println(moon["februari"]) //20

}
