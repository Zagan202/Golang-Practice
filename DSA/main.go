package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("-------------------------------------")
	//Demoing Searching Algorithms

	// fmt.Println("Demonstrating a linear search algorithm")
	// items := []int{12, 332, 42, 424, 23, 42, 90}
	// fmt.Println("Finding element 424 in int slice {12, 332, 42, 424, 23, 42, 90}")
	// fmt.Println(linearSearch(items, 424))
	// fmt.Println("Finding element 0 in int slice {12, 332, 42, 424, 23, 42, 90}")
	// fmt.Println(linearSearch(items, 0))
	// fmt.Println("-------------------------------------")
	// fmt.Println("Demonstrating a binary search algorithm")
	// items2 := []int{1, 2, 4, 7, 11}
	// fmt.Println("Finding element 542 in int slice {1, 32, 42, 124, 223, 542, 990}")
	// fmt.Println(binarySearch(items2, 7))
	// fmt.Println("Finding element 0 in int slice {1, 32, 42, 124, 223, 542, 990}")
	// fmt.Println(binarySearch(items2, 0))
	// fmt.Println("-------------------------------------")
	// fmt.Println("Demonstrating a interpolation search algorithm")
	// items3 := []int{1, 32, 42, 124, 223, 542, 990}
	// fmt.Println("Finding element 542 in int slice {1, 32, 42, 124, 223, 542, 990}")
	// fmt.Println(interpolationSearch(items3, 542))
	// fmt.Println("Finding element 0 in int slice {1, 32, 42, 124, 223, 542, 990}")
	// fmt.Println(interpolationSearch(items3, 0))

	//Demoing sorting Algorithms
	/*
		fmt.Println("Bubble Sort")
		slice := generateRandSlice(50)
		fmt.Println("\n--- Unsorted --- \n\n", slice)
		bubbleSort(slice)
		fmt.Println("\n--- Sorted ---\n\n", slice)
		fmt.Println("\nQuick Sort")
		slice2 := generateRandSlice(50)
		fmt.Println("\n--- Unsorted --- \n\n", slice2)
		quicksort(slice)
		fmt.Println("\n--- Sorted ---\n\n", slice2)
		fmt.Println("\nSelection Sort")
		slice3 := generateRandSlice(50)
		fmt.Println("\n--- Unsorted --- \n\n", slice3)
		selectionSort(slice3)
		fmt.Println("\n--- Sorted ---\n\n", slice3)
		fmt.Println("\nMerge Sort")
		slice4 := generateRandSlice(10)
		fmt.Println("\n--- Unsorted --- \n\n", slice4)
		fmt.Println("\n--- Sorted ---\n\n", mergeSort(slice4))
	*/
	//Demoing Binary tree
	tree := &BinaryTree{}
	tree.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)
	tree.insert(1000)
	tree.insert(0)
	print(os.Stdout, tree.root, 0, 'M')

	// //Demoing Stacks and Queues
	// s := NewStack()
	// s.Push(&Node{3})
	// s.Push(&Node{5})
	// s.Push(&Node{7})
	// s.Push(&Node{9})
	// fmt.Println(s.Pop(), s.Pop(), s.Pop(), s.Pop())

	// q := NewQueue(1)
	// q.Push(&Node{2})
	// q.Push(&Node{4})
	// q.Push(&Node{6})
	// q.Push(&Node{8})
	// fmt.Println(q.Pop(), q.Pop(), q.Pop(), q.Pop())

	// //Demoing Linked List

	// link := List{}
	// link.Insert("Hello")
	// link.Insert("This")
	// link.Insert("Is")
	// link.Insert("A")
	// link.Insert("Test")
	// link.Insert(link)

	// fmt.Println("\n==============================")
	// fmt.Printf("Head: %v\n", link.head.key)
	// fmt.Printf("Tail: %v\n", link.tail.key)
	// link.Display()
	// fmt.Println("\n==============================")
	// fmt.Printf("head: %v\n", link.head.key)
	// fmt.Printf("tail: %v\n", link.tail.key)
	// link.Reverse()
	// fmt.Println("\n==============================")
}
