package main

import (
	"fmt"
	"math/rand"
	"time"
	//"sort"
)

func main() {

	const PROBLEM_SIZE = 10000000
	const NUM_RUNS = 5

	var totalTime time.Duration

	anArray := [PROBLEM_SIZE]int {}

	for run := 0; run < NUM_RUNS; run++ {

		//fmt.Println(cap(anArray), len(anArray))

		for index := 0; index < len(anArray); index++ {
			anArray[index] = rand.Int()
		}

		before := time.Now()

		QuickSort(anArray[0:])

		fmt.Println("Run: ", run, "Duration: ", time.Now().Sub(before))
		totalTime += time.Now().Sub(before)

		//if sort.IntsAreSorted(anArray[0:]) {
		//	fmt.Println("Run 0:", run, "is sorted correctly.")
		//}
	}

	fmt.Println("The average duration for ", NUM_RUNS, " runs is: ", totalTime / NUM_RUNS)
}
// QuickSort entry point
// It initially checks for an empty or single-value slice to make sure we don't do any meaningless work
func QuickSort (unsorted []int) {

	if len(unsorted) > 1 {
		Partition(unsorted[0:])
	}


}

// Partition the passed in array
// 	For example, all values smaller than the pivot (the right-most element) are placed on the left side
// 	of the pivot, and all values greater than the pivot are placed on the right side of the pivot
func Partition (subSlice []int) []int {

	// Check if subSlice is empty or single-value, and return if so since we don't need to sort it
	if len(subSlice) < 2 {

		return subSlice

	// If the len of subSlice is 2, then swap the elements (if needed) and return
	} else if len(subSlice) == 2 {

		if subSlice[0] > subSlice[1] {
			subSlice[0], subSlice[1] = subSlice[1], subSlice[0]
		}

		return subSlice
	}

	// Get position of pivot (right-most element)
	pivotLoc := len(subSlice) - 1

	// Get position of leftLoc and rightLoc that will mark our comparison points (for swapping)
	leftLoc := 0
	rightLoc := pivotLoc - 1

	// isDone is used to exit out of the loop when the slice has been partitioned
	isDone := false

	for !isDone {

		// Find first element on the left that is greater than pivot, and check boundaries
		for (subSlice[leftLoc] < subSlice[pivotLoc]) && (leftLoc + 1 < len(subSlice)) {
			leftLoc++
		}

		// Find first element on the right that is less than pivot, and check boundaries
		for (subSlice[rightLoc] > subSlice[pivotLoc]) && (rightLoc - 1 > -1) {
			rightLoc--
		}

		// If leftLoc >= rightLoc, we have found where to swap our pivot in to be the partition separator
		if leftLoc >= rightLoc {
			subSlice[leftLoc], subSlice[pivotLoc] = subSlice[pivotLoc], subSlice[leftLoc]
			isDone = true // Mark that loop can be exited since subSlice has been partitioned

		// Else, we just swap the elements to their rightful side and continue looping
		} else if subSlice[leftLoc] != subSlice[rightLoc] {
			subSlice[leftLoc], subSlice[rightLoc] = subSlice[rightLoc], subSlice[leftLoc]
		}
	}

	// Recursively call Partition on left side of partition
	Partition(subSlice[0:leftLoc])

	// Recursively call partition on right side of partition
	Partition(subSlice[rightLoc:])

	return subSlice
}