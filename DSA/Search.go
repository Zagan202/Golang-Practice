package main

import "fmt"

func linearSearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

func binarySearch(haystack []int, needle int) bool {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2
		fmt.Println("median = ", median)
		if haystack[median] < needle {
			low = median + 1
			fmt.Println("low = ", low)
		} else {
			high = median - 1
			fmt.Println("high = ", high)
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}

func interpolationSearch(array []int, key int) int {

	min, max := array[0], array[len(array)-1]

	low, high := 0, len(array)-1

	for {
		if key < min {
			return low
		}

		if key > max {
			return high + 1
		}

		// make a guess of the location
		var guess int
		if high == low {
			guess = high
		} else {
			size := high - low
			offset := int(float64(size-1) * (float64(key-min) / float64(max-min)))
			guess = low + offset
		}

		// maybe we found it?
		if array[guess] == key {
			// scan backwards for start of value range
			for guess > 0 && array[guess-1] == key {
				guess--
			}
			return guess
		}

		// if we guessed to high, guess lower or vice versa
		if array[guess] > key {
			high = guess - 1
			max = array[high]
		} else {
			low = guess + 1
			min = array[low]
		}
	}
}
