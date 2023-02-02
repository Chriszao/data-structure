package main

import "fmt"

// Interpolation Search: an improvement over binary search best used for "uniformly" distributed data "guesses" where a value might be based on calculated probe results. If the probe is incorrect, the search area is narrowed, a new probe is narrowed, and a new probe is calculated.

// average case: O(log(log(n)))
// worst case: O(n) [values increase exponentially]

func interpolationSerach(array []int, target int) int {

	high := len(array) - 1
	low := 0
	steps := 1

	for target >= array[low] &&
		target <= array[high] &&
		low <= high {
		probe := low + (high-low)*(target-array[low])/(array[high]-array[low])

		fmt.Println("Probe:", probe)

		if array[probe] == target {
			fmt.Println("Took", steps, " steps to find the target")
			return probe
		} else if array[probe] < target {
			low = probe + 1
			steps++
		} else {
			high = probe - 1
			steps++
		}
	}

	fmt.Println("Took", steps, "and not found the target")

	return -1 // Target not found
}

func main() {
	array := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	if index := interpolationSerach(array, 256); index != -1 {
		fmt.Println("Found Index: ", index)
		return
	}

	fmt.Println("Element not found")
}
