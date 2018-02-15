package main

import (
	"fmt"
	"math"
)

func searchRecursive(a []int, low int, high int, key int) int {
	if high < low {
		return low
	}

	midFloat := float64(low + ((high - low) / 2))
	mid := int(math.Floor(midFloat))

	if key == a[mid] {
		return mid
	} else if key < a[mid] {
		return searchRecursive(a, low, mid-1, key)
	} else {
		return searchRecursive(a, mid+1, high, key)
	}
}

func searchIterative(a []int, low int, high int, key int) int {
	for low <= high {
		midFloat := float64(low + ((high - low) / 2))
		mid := int(math.Floor(midFloat))

		if key == a[mid] {
			return mid
		} else if key < a[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func main() {
	fmt.Println("Starting binary search")

	data := []int{3, 5, 8, 10, 12, 15, 18, 20, 20, 50, 60}

	result := searchRecursive(data, 0, 10, 50)
	fmt.Printf("Result resursive: %d \n", result)

	iterativeResult := searchIterative(data, 0, 10, 50)
	fmt.Printf("Result iterative: %d \n", iterativeResult)
}
