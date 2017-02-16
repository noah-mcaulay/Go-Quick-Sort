package main

import "fmt"

func main() {

	anArray := [10]int {10, 6, 7, 4, 15, 2, 3, 8, 9, 5}

	fmt.Println("Hello World!")

	fmt.Println(anArray)

	QuickSort(anArray[0:])

	fmt.Println(anArray)
}

func QuickSort (unsorted []int) []int {

	if len(unsorted) < 2 {
		return unsorted
	}

	//unsorted[0] = 42

	return Partition (unsorted[0:])
}

func Partition (subSlice []int) []int {

	pivotLoc := len(subSlice) - 1

	var test int = len(subSlice) - 1

	fmt.Print(test)
	leftLoc := 0
	rightLoc := pivotLoc - 1

	isDone := false

	for !isDone {

		for subSlice[leftLoc] < subSlice[pivotLoc] {
			leftLoc++
		}

		for subSlice[rightLoc] > subSlice[pivotLoc] {
			rightLoc--
		}

		if leftLoc >= rightLoc {
			subSlice[rightLoc], subSlice[pivotLoc] = subSlice[pivotLoc], subSlice[rightLoc]
			isDone = true
		} else if subSlice[leftLoc] != subSlice[rightLoc] {
			subSlice[leftLoc], subSlice[rightLoc] = subSlice[rightLoc], subSlice[leftLoc]
		}
	}



	//subSlice[pivotLoc], subSlice[leftLoc] = subSlice[rightLoc], subSlice[pivotLoc]


	return subSlice
}