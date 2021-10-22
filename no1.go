package main

import "fmt"

func primeNumber(n int) bool {

	if n > 1 {
		faktor := 0
		for i := 2; i <= n/2; i++ {
			if faktor == 0 {
				if n%i == 0 {
					faktor++
				}
			} else {
				return false
			}
		}
		return true
	} else {
		return false
	}

}

func main() {
	fmt.Println("Result :")
	fmt.Println("1000000007 : ", primeNumber(1000000007))   // true
	fmt.Println("1500450271 : ", primeNumber(1500450271))   // true
	fmt.Println("1000000000 : ", primeNumber(1000000000))   // false
	fmt.Println("10000000019 : ", primeNumber(10000000019)) // true
	fmt.Println("10000000033 : ", primeNumber(10000000033)) // true
}
