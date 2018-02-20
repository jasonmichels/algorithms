package insertion_sort

// GenerateData Generate an array that is backwards to test worst case performance
func GenerateData(max int) []int {
	data := make([]int, max+1)

	currentNum := max

	for i := 0; i <= max; i++ {
		data[i] = currentNum

		if currentNum > 0 {
			currentNum = currentNum - 1
		}
	}

	return data
}

// Implementation
// data := is.GenerateData(100)
// 	result := is.Naive(data)
// 	fmt.Println(len(result))

func Naive(data []int) []int {
	// loop through all elements starting with index 1
	for i := 1; i < len(data); i++ {
		var temp int

		if data[i] < data[i-1] {
			// check to make sure the element is bigger than the next element, and keep switching until that is not true
			for j := i; j > 0; j-- {
				if data[j] < data[j-1] {
					temp = data[j-1]
					data[j-1] = data[j]
					data[j] = temp
				} else {
					break
				}
			}

		}
	}

	return data
}
