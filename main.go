package main

import (
	"fmt"

	"github.com/jasonmichels/algorithms/peak"
)

func main() {
	nums := peak.GenerateLargeNumSlice(8000000, 10000000)
	// index := peak.PeakNaive(nums)
	// index := peak.PeakDivide(nums)
	index := peak.PeakDivideConquer(nums)
	fmt.Printf("Return value: %v \n", index)
}
