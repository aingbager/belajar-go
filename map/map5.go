package main

import "fmt"

func main() {
	//belajar map

	var bulan map[string]int
	bulan = map[string]int{}
	bulan["januari"] = 1
	bulan["februari"] = 2
	bulan["maret"] = 3

	fmt.Println(bulan["januari"])
	fmt.Println(bulan["februari"])
}
