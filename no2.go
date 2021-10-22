package main

import (
	"fmt"
	"time"
)

func timeMeasurement(start time.Time) {
	waktu := time.Since(start)
	fmt.Printf("Execution time: %s", waktu)
}

func pow(x, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * x
		}
		n = n / 2
		x = x * x
	}
	return res
}
func main() {
	defer timeMeasurement(time.Now())
	fmt.Println("Result =")
	fmt.Println(pow(5, 2))
	fmt.Println(pow(2, 3))
	fmt.Println(pow(7, 2))
	fmt.Println(pow(10, 5))
	fmt.Println(pow(17, 6))
	fmt.Println(pow(5, 3))
}
