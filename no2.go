package main

import "fmt"

func fibo(n int) int {
	x := make([]int, n+2)
	x[0] = 0
	x[1] = 1
	for i := 2; i < n+1; i++ {
		x[i] = x[i-1] + x[i-2]
	}
	return x[n]
}

func main() {
	fmt.Println(fibo(0))
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(3))
	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))
	fmt.Println(fibo(9))
	fmt.Println(fibo(10))
}
