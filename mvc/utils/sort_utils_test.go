package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSortWorstCase(t *testing.T) {
	// Initialization
	els := []int{9, 8, 7, 2, 5}

	// Execution
	BubbleSort(els)

	// Validation
	assert.NotNil(t, els)
	assert.EqualValuesf(t, 5, len(els), "")
	assert.EqualValuesf(t, []int{2, 5, 7, 8, 9}, els, "")
}

func TestBubbleSortBestCase(t *testing.T) {
	// Initialization
	els := []int{2, 5, 7, 8, 9}

	// Execution
	BubbleSort(els)

	// Validation
	assert.NotNil(t, els)
	assert.EqualValuesf(t, 5, len(els), "")
	assert.EqualValuesf(t, []int{2, 5, 7, 8, 9}, els, "")
}

func TestBubbleSortNilSlice(t *testing.T) {
	// Initialization

	// Execution
	BubbleSort(nil)

	// Validation

}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	els := getElements(5)
	assert.EqualValuesf(t, 5, len(els), "")
	assert.EqualValuesf(t, []int{4, 3, 2, 1, 0}, els, "")
}

func BenchmarkBubbleSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		//BubbleSort(els)
		Sort(els)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		//BubbleSort(els)
		Sort(els)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		//BubbleSort(els)
		Sort(els)
	}
}

func BenchmarkSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkSort1000(b *testing.B) {
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}
