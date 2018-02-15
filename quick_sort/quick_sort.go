package main

import (
	"fmt"
)

func Partition(a []int, low int, high int) int {
	i := low // 0
	j := high // 16

	tmp := 0

	pivot := a[(low + high) / 2] // 5

	//8, 4, 2, 5, 2, 100, 23, 4, 5, 6, 8, 1, 2, 3, 4, 98, 4

	for i <= j {
		for a[i] < pivot {
			i++
		}

		for a[j] > pivot {
			j--
		}

		if i <= j {

			tmp = a[i]
			a[i] = a[j]
			a[j] = tmp
			i++
			j--
		}
	}

	fmt.Println(a)

	return i
}

func QuickSort(a []int, low int, high int) []int {

	index := Partition(a, low, high)


	if low < index - 1 {
		QuickSort(a, low, index - 1)
	}

	if index < high {
		QuickSort(a, index, high)
	}

	return a

}

func main()  {
	fmt.Println("Quick sort algorithm")

	n := []int{8, 4, 2, 5, 2, 100, 23, 4, 5, 6, 8, 1, 2, 3, 4, 98, 4}
	result := QuickSort(n, 0, len(n) - 1)

	fmt.Println(result)
}
