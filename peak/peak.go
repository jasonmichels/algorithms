package peak

func GenerateLargeNumSlice(peak, largest int) []int {
	nums := make([]int, largest, largest)

	for i := 0; i <= peak; i++ {
		nums[i] = i
	}

	for i := peak + 1; i < largest; i++ {
		nums[i] = nums[i-1] - 1
	}
	return nums
}

func PeakNaive(nums []int) int {
	var response int

	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			response = nums[i]
			break
		}

		if nums[i] >= nums[i+1] {
			response = nums[i]
			break
		}
	}

	return response
}

func PeakDivide(nums []int) int {
	var response int
	half := len(nums) / 2

	if nums[half] >= nums[half-1] && nums[half] >= nums[half+1] {
		response = nums[half]
		return response
	}

	if nums[half] < nums[half-1] {
		return PeakNaive(nums[0:half])
	} else if nums[half] < nums[half+1] {
		return PeakNaive(nums[half:len(nums)])
	}

	return response
}

func FindIndex(nums []int) (int, bool) {
	half := len(nums) / 2

	if len(nums) == 2 {
		return nums[1], true
	}

	if nums[half] >= nums[half-1] && nums[half] >= nums[half+1] {
		return nums[half], true
	}
	return 0, false
}

func PeakDivideConquer(nums []int) int {
	var response int
	looking := true

	for looking {
		if index, found := FindIndex(nums); found == true {
			response = index
			looking = false
			break
		}

		half := len(nums) / 2

		if nums[half] < nums[half-1] {
			nums = nums[0:half]
		} else if nums[half] < nums[half+1] {
			nums = nums[half:len(nums)]
		}

		if len(nums) == 0 {
			break
		}
	}

	return response
}
