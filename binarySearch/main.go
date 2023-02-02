package main

import "fmt"

func binarySearch(array []int, target int) (index int) {
	low := 0
	high := len(array) - 1
	steps := 1

	for low <= high {
		middle := low + (high-low)/2
		value := array[middle]

		fmt.Println("Step: ", steps, " Middle: ", value)

		if value < target {
			low = middle + 1
			steps++
		} else if value > target {
			high = middle - 1
			steps++
		} else {
			return middle
		}
	}

	return -1 // Target not found
}

func main() {
	array := make([]int, 1000000)
	target := 777777

	for index := 0; index < len(array); index++ {
		array[index] = index
	}

	index := binarySearch(array, target)
	fmt.Println("Founded index:", index)
}
