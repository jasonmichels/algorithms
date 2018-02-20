package insertion_sort

import (
	"testing"
)

func TestGenerateData(t *testing.T) {
	max := 50000
	nums := GenerateData(max)

	for i := 0; i <= max; i++ {
		if nums[i] != max {
			t.Errorf("Our data was not generated correctly at index: %v", i)
		}

		if max > 0 {
			max = max - 1
		}
	}
}

func BenchmarkGenerateData(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateData(50000)
	}
}

func TestNaive(t *testing.T) {
	max := 35000
	nums := GenerateData(max)

	result := Naive(nums)

	for i := 0; i <= max; i++ {
		if result[i] != i {
			t.Errorf("Our data is not sorted correctly at index: %v", i)
		}
	}
}

func BenchmarkNaive(b *testing.B) {
	max := 35000
	nums := GenerateData(max)

	for n := 0; n < b.N; n++ {
		Naive(nums)
	}
}
