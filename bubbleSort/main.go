package main

import "fmt"

// BubbleSort: pairs of adjacent elemnts are compared, and the elements swapped if they are not in order.

// Time Complexity = Quadratic time O(n^2)
// It's okay to use with small data set, but not recommended to use with large data set

func bubbleSort(array []int) {
	for index := 0; index < len(array)-1; index++ {
		for nestIndex := 0; nestIndex < len(array)-index-1; nestIndex++ {
			if array[nestIndex] > array[nestIndex+1] {
				array[nestIndex], array[nestIndex+1] = array[nestIndex+1], array[nestIndex]
			}
		}
	}
}

func main() {
	array := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}

	bubbleSort(array)

	for _, value := range array {
		fmt.Print(value)
	}
}
