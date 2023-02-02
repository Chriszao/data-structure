package main

import "fmt"

// Selection sort: Search through an array and keep track of the minimum value during each iteration. At the end of each iteration, we swap variables.

// Time complexity = Quadratic time O(n^2)
// It's okay to use with small data set, but not recommended to use with large data set

func selectionSort(array []int) {
	for index := 0; index < len(array)-1; index++ {
		min := index

		for nestIndex := index; nestIndex < len(array); nestIndex++ {
			if array[min] > array[nestIndex] {
				min = nestIndex
			}
		}

		array[index], array[min] = array[min], array[index]
	}
}

func main() {
	array := []int{8, 7, 9, 2, 3, 1, 5, 4, 6}

	selectionSort(array)

	for _, value := range array {
		fmt.Print(value)
	}
}
