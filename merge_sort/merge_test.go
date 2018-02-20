package merge_sort

import (
	"testing"

	is "github.com/jasonmichels/algorithms/insertion_sort"
)

func TestNaive(t *testing.T) {
	max := 35000
	nums := is.GenerateData(max)

	result := Naive(nums)

	for i := 0; i <= max; i++ {
		if result[i] != i {
			t.Errorf("Our data is not sorted correctly at index: %v", i)
		}
	}
}

func BenchmarkNaive(b *testing.B) {
	max := 35000
	nums := is.GenerateData(max)

	for n := 0; n < b.N; n++ {
		Naive(nums)
	}
}

func TestSecond(t *testing.T) {
	max := 35000
	nums := is.GenerateData(max)

	result := Second(nums)

	for i := 0; i <= max; i++ {
		if result[i] != i {
			t.Errorf("Our data is not sorted correctly at index: %v", i)
		}
	}
}

func BenchmarkSecond(b *testing.B) {
	max := 35000
	nums := is.GenerateData(max)

	for n := 0; n < b.N; n++ {
		Second(nums)
	}
}
