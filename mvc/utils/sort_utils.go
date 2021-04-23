package utils

import "sort"

// []int {9, 8, 7, 5, 2, 6}
// expected output []int{2, 5, 6, 7, 8, 9}

func BubbleSort(elements []int) {
	keepRunning := true
	for keepRunning {
		keepRunning = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}
}

func Sort(els []int) {
	if len(els) < 1000 {
		BubbleSort(els)
	} else {
		sort.Ints(els)
	}
}
