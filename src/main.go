package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	"runtime"
	"sync"
)

var wg sync.WaitGroup


func main() {

	anArray := [20]int {}

	for i := 0; i < len(anArray); i++ {
		anArray[i] = rand.Int()
	}

	fmt.Println("Hello World!")

	//fmt.Println(anArray)


	fmt.Println(anArray)

	before := time.Now()

	runtime.GOMAXPROCS(2)
	wg.Add(2)
	go QuickSort(anArray[0:len(anArray)/2-1])
	go QuickSort(anArray[len(anArray)/2:])
	wg.Wait()

	fmt.Println(time.Now().Sub(before))

	if sort.IntsAreSorted(anArray[0:]) {
		fmt.Println("We did it bois")
	}

	fmt.Println(anArray)
}

func QuickSort (unsorted []int) {

	//arrayLen := len(unsorted)

	if len(unsorted) < 2 {
		return
	}

	//unsorted[0] = 42

	Partition (unsorted[0:])
	//for i := 0; i < numProcs; i++ {
	//	go Partition
	//}

	defer wg.Done()
}

func Partition (subSlice []int) []int {

	if len(subSlice) < 2 {

		return subSlice

	} else if len(subSlice) == 2 {

		if subSlice[0] > subSlice[1] {
			subSlice[0], subSlice[1] = subSlice[1], subSlice[0]
		}

		return subSlice
	}

	pivotLoc := len(subSlice) - 1

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
			subSlice[leftLoc], subSlice[pivotLoc] = subSlice[pivotLoc], subSlice[leftLoc]
			isDone = true
		} else if subSlice[leftLoc] != subSlice[rightLoc] {
			subSlice[leftLoc], subSlice[rightLoc] = subSlice[rightLoc], subSlice[leftLoc]
		}
	}

	Partition(subSlice[0:leftLoc])
	Partition(subSlice[rightLoc:])

	return subSlice
}