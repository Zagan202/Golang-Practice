package main

import (
	"fmt"
)

func main() {
	fmt.Println("-------------------------------------")
	//Demoing Searching Algorithms
	/*
		fmt.Println("Demonstrating a linear search algorithm")
		items := []int{12, 332, 42, 424, 23, 42, 90}
		fmt.Println("Finding element 424 in int slice {12, 332, 42, 424, 23, 42, 90}")
		fmt.Println(linearSearch(items, 424))
		fmt.Println("Finding element 0 in int slice {12, 332, 42, 424, 23, 42, 90}")
		fmt.Println(linearSearch(items, 0))
		fmt.Println("-------------------------------------")
		fmt.Println("Demonstrating a binary search algorithm")
		items2 := []int{1, 32, 42, 124, 223, 542, 990}
		fmt.Println("Finding element 542 in int slice {1, 32, 42, 124, 223, 542, 990}")
		fmt.Println(binarySearch(items2, 542))
		fmt.Println("Finding element 0 in int slice {1, 32, 42, 124, 223, 542, 990}")
		fmt.Println(binarySearch(items2, 0))
		fmt.Println("-------------------------------------")
		fmt.Println("Demonstrating a interpolation search algorithm")
		items3 := []int{1, 32, 42, 124, 223, 542, 990}
		fmt.Println("Finding element 542 in int slice {1, 32, 42, 124, 223, 542, 990}")
		fmt.Println(interpolationSearch(items3, 542))
		fmt.Println("Finding element 0 in int slice {1, 32, 42, 124, 223, 542, 990}")
		fmt.Println(interpolationSearch(items3, 0))
	*/
	//Demoing sorting Algorithms
	fmt.Println("BubbleSort")
	slice := generateRandSlice(100)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	bubbleSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice)
	fmt.Println("\nQuickSort")
	slice2 := generateRandSlice(100)
	fmt.Println("\n--- Unsorted --- \n\n", slice2)
	quicksort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice2)
}
