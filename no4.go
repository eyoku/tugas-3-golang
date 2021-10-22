package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	var outSekali []int
	var validator = make(map[rune]int)
	for _, val := range angka {
		validator[val] = validator[val] + 1
	}
	for number := range validator {
		if validator[number] == 1 {
			outSekali = append(outSekali, int(number)-48)
		}
	}
	return outSekali
}

func main() {
	fmt.Println("1234123 : ", munculSekali("1234123"))
	fmt.Println("76523752 : ", munculSekali("76523752"))
	fmt.Println("12345 : ", munculSekali("12345"))
	fmt.Println("1122334455 : ", munculSekali("1122334455"))
	fmt.Println("0872504 : ", munculSekali("0872504"))
}
