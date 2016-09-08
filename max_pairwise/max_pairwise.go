package main

import "fmt"

func maxPairwiseProductFast(numbers []int) int {

	var maxIndexOne int = -1

	for index, value := range numbers {
		if maxIndexOne == -1 || (value > numbers[maxIndexOne]) {
			maxIndexOne = index
		}
	}

	var maxIndexTwo int = -1

	for index, value := range numbers {
		if ((maxIndexTwo == -1) || (value >= numbers[maxIndexTwo])) && (index != maxIndexOne) {
			maxIndexTwo = index
		}
	}

	return numbers[maxIndexOne] * numbers[maxIndexTwo]
}

func main() {
	numbers := []int{1, 2, 3, 3, 4, 4}

	result := maxPairwiseProductFast(numbers)

	fmt.Println(result)
}
