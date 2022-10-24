package sorting

import (
	"testing"
)

var (
	caso  = []int{1, 2, 46, 7, 3, 2, 6, 6, 7 - 5}
	caso1 = []int{1: 10, 9999: 10}
)

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(caso)
	}
}
func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort(caso)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(caso, 0, len(caso)-1)
	}
}
