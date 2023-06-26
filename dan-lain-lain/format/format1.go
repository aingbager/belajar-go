package main

import (
	"fmt"
)

type student struct {
	name       string
	height     float64
	age        int32
	isGraduate bool
	hobbies    []string
}

var data = student{
	name:       "whick",
	height:     182.2,
	age:        26,
	isGraduate: false,
	hobbies:    []string{"makan", "minum"},
}

func main() {

	fmt.Printf("%b\n", data.age) //type data %b kenversi data jadi numerik biner

	fmt.Printf("%d\n", data.age)    //type data %d mempormat data numerik jadi string numerik
	fmt.Printf("%e\n", data.height) // kovnersi data float ke notasi numerik
	// 1.825000e+02
	fmt.Printf("%E\n", data.height)
	// 1.825000E+02
}
