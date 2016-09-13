package main

import (
	"fmt"
	"math"
)

func gcdSlow(a int, b int) int {
	var best int = 0

	for i := 1; i <= a + b; i++ {

		if math.Mod(float64(a), float64(i)) == 0 && math.Mod(float64(b), float64(i)) == 0 {		
			best = i
		}
	}

	return best
}

func gcdFast(a int, b int) int{
	// fmt.Printf("%d ---- %d \n", a, b)
	if b == 0 {
		return a
	} else {
		return gcdFast(b, int(math.Mod(float64(a), float64(b))))
	}
}

func main() {
	// fmt.Printf("GCD slow: %d \n", gcdSlow(3918848, 1653264))
	fmt.Printf("GCD fast: %d \n", gcdFast(357, 234))
}