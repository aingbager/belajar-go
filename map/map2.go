package main

import "fmt"

func main() {
	var bulan = map[string]int{
		"januari":  50,
		"february": 40,
	}

	fmt.Println(bulan["januari"])
	fmt.Println(bulan["february"])

	//mencetak dengan menggunakan for,tambahkan range di depan variable

	for key, val := range bulan {
		fmt.Println(key, " \t:", val)
	}
}
