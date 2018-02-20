package main

import (
	"fmt"

	is "github.com/jasonmichels/algorithms/insertion_sort"
	"github.com/jasonmichels/algorithms/merge_sort"
)

func main() {
	data := is.GenerateData(10)
	result := merge_sort.Second(data)
	fmt.Println(len(result))
	fmt.Println(result)
}
