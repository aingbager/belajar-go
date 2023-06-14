// belajar fungsi
package main

import (
	"fmt"
	"strings"
)

func hasil(name string, names []string) {
	name1 := strings.Join(names, " ")
	fmt.Println(name, name1)
}

func main() {
	names := []string{"otong", "surotong"}
	hasil("hello", names)
}
