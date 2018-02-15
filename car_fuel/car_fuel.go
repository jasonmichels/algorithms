package main

import "fmt"

func MinRefills(x []int, n int, L int) int {
	numRefills := 0
	currentRefill := 0
	for currentRefill <= n {
		lastRefill := currentRefill

		// implement

		if currentRefill == lastRefill {
			fmt.Println("Some mistake made")
		}
		if currentRefill <= n {
			numRefills = numRefills + 1
			fmt.Println("Number of refills %i", numRefills)
		}
	}
	return numRefills
}
func main() {
	x := []int{0, 100, 350, 600, 700, 900}
	n := 5
	L := 400

	MinRefills(x, n, L)
}
