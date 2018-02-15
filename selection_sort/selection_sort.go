package main

import (
	"fmt"
)

func selectionSort(n []int) []int  {
	minIndex := 0
	for i := range n {
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[minIndex] {
				minIndex = j
			}
		}
		prev := n[i]
		min := n[minIndex]

		n[i] = min
		n[minIndex] = prev

		minIndex = i + 1
	}

	return n
}

func MergeSort(n []int) []int  {
	if len(n) <= 1 {
		return n
	}

	mid := len(n) / 2
	left := MergeSort(n[:mid])
	right := MergeSort(n[mid:])

	return Merge(left, right)
}

func Merge(left, right []int) []int{
	fmt.Println(left)
	fmt.Println(right)

	leftIndex := 0
	rightIndex := 0

	size := len(left) + len(right)
	result := make([]int, size, size)

	for index := range result {

		if len(left) == leftIndex && len(right) > rightIndex {
			result[index] = right[rightIndex]
			rightIndex++
		} else if len(right) == rightIndex && len(left) > leftIndex{
			result[index] = left[leftIndex]
			leftIndex++
		} else if left[leftIndex] < right[rightIndex] {
			result[index] = left[leftIndex]
			leftIndex++
		} else {
			result[index] = right[rightIndex]
			rightIndex++
		}
	}

	fmt.Println(result)
	fmt.Println("------------")

	return result
}

func main()  {
	fmt.Println("Selection sort algorithm")

	n := []int{8, 4, 2, 5, 2, 100, 23, 4, 5, 6, 8, 1, 2, 3, 4, 98, 99}
	//fmt.Println(selectionSort(n))
	fmt.Println(MergeSort(n))

}