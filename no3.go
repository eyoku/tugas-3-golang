package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	var xs []int
	temp := int(math.Abs(float64(jumps[0] - jumps[1])))
	xs = append(xs, 0)
	xs = append(xs, temp)
	for i := 2; i < len(jumps); i++ {
		temp_1 := xs[i-1] + int(math.Abs(float64(jumps[i]-jumps[i-1])))
		temp_2 := xs[i-2] + int(math.Abs(float64(jumps[i]-jumps[i-2])))
		if temp_1 < temp_2 {
			xs = append(xs, temp_1)
		} else {
			xs = append(xs, temp_2)
		}
	}
	return xs[len(jumps)-1]

}
func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))
}
