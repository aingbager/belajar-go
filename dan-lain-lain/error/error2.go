package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("tidak boleh kosong")
	}
	return true, nil
}

func main() {
	//belajar error
	var name string
	fmt.Print("your name:")
	fmt.Sprint(&name)

	if valid, err := validate(name); valid {
		fmt.Println("hallo", name)
	} else {
		fmt.Println(err.Error())
	}
}
