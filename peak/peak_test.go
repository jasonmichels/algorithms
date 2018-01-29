package peak

import (
	"testing"
)

func TestPeakNaive(t *testing.T) {
	nums := GenerateLargeNumSlice(8000000, 10000000)

	if index := PeakNaive(nums); index != 8000000 {
		t.Errorf("We received incorrect index %v", index)
	}
}

func BenchmarkPeakNaive(b *testing.B) {
	nums := GenerateLargeNumSlice(8000000, 10000000)

	for n := 0; n < b.N; n++ {
		PeakNaive(nums)
	}
}

func BenchmarkGenerateLargeNumSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateLargeNumSlice(8000000, 10000000)
	}
}

func TestPeakDivide(t *testing.T) {
	nums := GenerateLargeNumSlice(8000000, 10000000)

	if index := PeakDivide(nums); index != 8000000 {
		t.Errorf("We received incorrect index %v", index)
	}
}

func BenchmarkPeakDivide(b *testing.B) {
	nums := GenerateLargeNumSlice(8000000, 10000000)

	for n := 0; n < b.N; n++ {
		PeakDivide(nums)
	}
}

func TestPeakPeakDivideConquer(t *testing.T) {
	nums := GenerateLargeNumSlice(8000000, 10000000)

	if index := PeakDivideConquer(nums); index != 8000000 {
		t.Errorf("We received incorrect index %v", index)
	}
}

func BenchmarkPeakDivideConquer(b *testing.B) {
	nums := GenerateLargeNumSlice(8000000, 10000000)

	for n := 0; n < b.N; n++ {
		PeakDivideConquer(nums)
	}
}
