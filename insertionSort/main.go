package main

import "fmt"

// Insertion Sort: After comparing to the left shift elements to the right to make room to insert a value

// Time complexity = Quadratic time O(n^2)
// It's okay to use with small data set, but not recommended to use with large data set

// Less steps than bubble sort
// Best case is O(n) compared to selection Sort O(n^2)

func insertionSort(array []int) {
	for index := 1; index < len(array); index++ {
		temp := array[index]
		previousItemIndex := index - 1

		for previousItemIndex >= 0 && array[previousItemIndex] > temp {
			array[previousItemIndex+1] = array[previousItemIndex]
			previousItemIndex--
		}

		array[previousItemIndex+1] = temp
	}
}

func main() {
	array := []int{9, 1, 8, 2, 7, 3, 6, 5, 4}

	insertionSort(array)

	for _, value := range array {
		fmt.Print(value, " ")
	}
}
