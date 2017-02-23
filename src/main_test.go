package main

import (
	"testing"
	"math/rand"
	//"sort"
	"sort"
)

// TODO: Switch to table driven tests if possible.
func TestQuickSort(t *testing.T) {

	rand.Seed(42)

	Test1K := [1000]int{}
	Test5  := [5]int{10, 7, 1, 3, 4}
	Test4  := [4]int{2, 10, 3, 20}
	Test3  := [3]int{1, 6, 3}
	Test2  := [2]int{5, 2}
	Test1  := [1]int{10}
	Test0  := [0]int{}

	// Populate Test1K with random values
	for i := 0; i < len(Test1K); i++ {
		Test1K[i] = rand.Int()
	}

	// ---- Get sorted arrays ----
	QuickSort(Test1K[0:])
	QuickSort(Test5[0:])
	QuickSort(Test4[0:])
	QuickSort(Test3[0:])
	QuickSort(Test2[0:])
	QuickSort(Test1[0:])
	QuickSort(Test0[0:])

	// ---- Test cases ----
	if !sort.IntsAreSorted(Test1K[0:]) {
		t.Error("Failed to sort 1K elements.")
	}

	if !sort.IntsAreSorted(Test5[0:]) {
		t.Error("Failed to sort 5 elements.")
	}

	if !sort.IntsAreSorted(Test4[0:]) {
		t.Error("Failed to sort 3 elements.")
	}

	if !sort.IntsAreSorted(Test2[0:]) {
		t.Error("Failed to sort 2 elements.")
	}

	if !sort.IntsAreSorted(Test1[0:]) {
		t.Error("Failed to sort 1 element.")
	}

	if !sort.IntsAreSorted(Test0[0:]) {
		t.Error("Failed to sort 0 elements.")
	}


}

func BenchmarkQuickSort1M(b *testing.B) {

	rand.Seed(42)

	Bench1M := [10000]int{}

	// Populate Test1K with random values
	for i := 0; i < len(Bench1M[0:]); i++ {
		Bench1M[i] = rand.Int()
	}

	b.ResetTimer()

	// ---- Get sorted arrays ----
	for n := 0; n < b.N; n++ {
		QuickSort(Bench1M[0:])
	}


}

func BenchmarkQuickSort9K(b *testing.B) {

	rand.Seed(42)

	Bench9K := [9000]int{}

	// Populate Test9K with random values
	for i := 0; i < len(Bench9K[0:]); i++ {
		Bench9K[i] = rand.Int()
	}

	// ---- Get sorted arrays ----
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		QuickSort(Bench9K[0:])
	}


}