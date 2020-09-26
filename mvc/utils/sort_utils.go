package utils

import "sort"

// []int {9,8,7,6,5}
// []int {5,6,7,8,9}
func BubbleSort(elements []int) []int {
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
	return elements
}

func Sort(els []int) {
	// NOTE: Sort function is faster than native sort and bubble sort alone
	// benchmark helps us to check if the current algorithm we have is the efficient one, i.e. to check for performance
	if len(els) < 1000 {
		BubbleSort(els)
		return
	}
	sort.Ints(els)
}
