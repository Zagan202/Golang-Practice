package main

import "fmt"

// Driver program to test above Huffman Coding
func main() {
	test := "abcdefghijklmnopqrstuvwxyz"

	symFreqs := make(map[rune]int)
	// read each symbol and record the frequencies
	for _, c := range test {
		symFreqs[c]++
	}

	// example tree
	exampleTree := buildTree(symFreqs)

	// print out results
	fmt.Println("SYMBOL\tWEIGHT\tHUFFMAN CODE")
	printCodes(exampleTree, []byte{})
}
