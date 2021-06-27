package utils

import (
	"sort"
	"testing"
)

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestBobbleSortWorstCase(t *testing.T) {
	els := []int{9, 8, 7, 6, 5}
	els = BobbleSort(els)

	if els == nil {
		t.Error("slice should not be nil")
	}
	if els != nil && len(els) != 5 {
		t.Error("slice should have lent of 5")
	}
	if !sort.IntsAreSorted(els) {
		t.Error("slice should be sorted")
	}

}

func TestBobbleSortBestCase(t *testing.T) {
	els := []int{5, 6, 7, 8, 9}
	els = BobbleSort(els)

	if els == nil {
		t.Error("slice should not be nil")
	}
	if els != nil && len(els) != 5 {
		t.Error("slice should have lent of 5")
	}
	if !sort.IntsAreSorted(els) {
		t.Error("slice should be sorted", els)
	}
}

func BenchmarkBobbleSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		BobbleSort(els)
	}
}

func BenchmarkSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkBobbleSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		BobbleSort(els)
	}
}

func BenchmarkSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}
