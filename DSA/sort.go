package main

import (
	"log"
	"math/rand"
	"time"
)

//Time tracking trick taken from https://blog.stathat.com/2012/10/10/time_any_function_in_go.html
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func generateRandSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func bubbleSort(items []int) {

	var (
		n      = len(items)
		sorted = false
	)

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}

func quicksort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	left, right := 0, len(items)-1

	pivot := rand.Int() % len(items)

	items[pivot], items[right] = items[right], items[pivot]

	for i := range items {
		if items[i] < items[right] {
			items[left], items[i] = items[i], items[left]
			left++
		}
	}

	items[left], items[right] = items[right], items[left]

	quicksort(items[:left])
	quicksort(items[left+1:])

	return items
}

func selectionsort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}
