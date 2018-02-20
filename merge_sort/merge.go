package merge_sort

func Naive(data []int) []int {
	// 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0

	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2

	left := Naive(data[:mid])
	right := Naive(data[mid:])

	return Merge(left, right)
}

func Merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

func Second(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2

	left := Second(data[:mid])
	right := Second(data[mid:])

	return MergeSecond(left, right)
}

func MergeSecond(left, right []int) []int {
	size := len(left) + len(right)
	var result []int

	for i := 0; i < size; i++ {
		if (len(left) > 0) && (len(right) > 0) && (left[0] < right[0]) {
			result = append(result, left[0])
			if len(left) > 1 {
				left = left[1:]
			} else {
				left = []int{}
			}

		} else if (len(right) > 0) && (len(left) > 0) && (right[0] < left[0]) {
			result = append(result, right[0])
			if len(right) > 1 {
				right = right[1:]
			} else {
				right = []int{}
			}
		} else if len(left) > 0 && len(right) == 0 {
			result = append(result, left[0])
			if len(left) > 1 {
				left = left[1:]
			} else {
				left = []int{}
			}
		} else if len(right) > 0 && len(left) == 0 {
			result = append(result, right[0])
			if len(right) > 1 {
				right = right[1:]
			} else {
				right = []int{}
			}
		}
	}

	return result
}
