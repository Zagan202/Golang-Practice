package main

import (
	"fmt"
)

func sliceStuff() {
	fmt.Println("Type: basic \nA slice can be initialized at runtime using the built-in function make")
	fmt.Println("Type: new dec \nA slice can be Declaration using new keyword")
	fmt.Println("Type: literal \nLiteral declaration []slice_type{comma-separated list of element values}")
	fmt.Println("Type: append \nLiteral declaration []slice_type{comma-separated list of element values}")
	fmt.Println("Type: copy \nEnlarge a slice using the copy function" +
		"\nGo's built-in copy function is used to copy data from one slice to another." +
		"\ncopy takes two arguments: dst and src. All of the entries in src are copied into dst overwriting whatever is there." +
		"\nIf the lengths of the two slices are not the same, the smaller of the two will be used")
	fmt.Println("Type: tricks \nvarious slicing tricks")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "basic":
		var intSlice1 = make([]int, 10)     // when length and capacity is same
		var intSlice2 = make([]int, 10, 20) // when length and capacity is different

		fmt.Printf("intSlice1 \tLen: %v \tCap: %v\n", len(intSlice1), cap(intSlice1))
		fmt.Printf("intSlice1 \tLen: %v \tCap: %v\n", len(intSlice2), cap(intSlice2))
	case "new dec":
		var intSlice1 = new([50]int)[0:10]
		fmt.Printf("intSlice1 \tLen: %v \tCap: %v\n", len(intSlice1), cap(intSlice1))
	case "literal":
		var intSlice = []int{10, 20, 30, 40}
		var strSlice = []string{"India", "Canada", "Japan"}
		fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
		fmt.Println(intSlice)
		fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
		fmt.Println(strSlice)
	case "append":
		// Create a smaller slice
		a := make([]int, 2, 5)
		a[0] = 10
		a[1] = 20
		fmt.Println("Slice A:", a)
		fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))

		// Create a bigger slice
		a = append(a, 30, 40, 50, 60, 70, 80, 90)
		fmt.Println("Slice A after appending data:", a)
		fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))
	case "copy":
		// Create a smaller slice
		a := []int{5, 6, 7}
		fmt.Printf("[Slice:A] Length is %d Capacity is %d\n", len(a), cap(a))
		// Create a bigger slice
		b := make([]int, 5, 10)
		copy(b, a) // copy function
		fmt.Printf("[Slice:B] Length is %d Capacity is %d\n", len(b), cap(b))

		fmt.Println("Slice B after copying:", b)
		b[3] = 8
		b[4] = 9
		fmt.Println("Slice B after adding elements:", b)
	case "tricks":
		var countries = []string{"india", "japan", "canada", "australia", "russia"}

		fmt.Printf("Countries: %v\n", countries)

		fmt.Printf(":2 %v\n", countries[:2])

		fmt.Printf("1:3 %v\n", countries[1:3])

		fmt.Printf("2: %v\n", countries[2:])

		fmt.Printf("2:5 %v\n", countries[2:5])

		fmt.Printf("0:3 %v\n", countries[0:3])

		fmt.Printf("Last element: %v\n", countries[4])
		fmt.Printf("Last element: %v\n", countries[len(countries)-1])
		fmt.Printf("Last element: %v\n", countries[4:])

		fmt.Printf("All elements: %v\n", countries[0:len(countries)])

		fmt.Printf("Last two elements: %v\n", countries[3:len(countries)])
		fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)])

		fmt.Println(countries[:])
		fmt.Println(countries[0:])
		fmt.Println(countries[0:len(countries)])
	case "slice to slice":
		var oldStr = []string{"india", "japan", "canada", "australia", "russia"}
		var strSlice []string
		newStr := oldStr[0:3]
		strSlice = append(strSlice, oldStr[:1]...)
		fmt.Printf("newStr: %v\n", newStr)
		fmt.Printf("strSlice: %v\n", strSlice)

		fmt.Printf("oldStr length: %v\tcapacity: %v\n", len(oldStr), cap(oldStr))
		fmt.Printf("newStr length: %v\tcapacity: %v\n", len(newStr), cap(newStr))
		fmt.Printf("strSlice length: %v\tcapacity: %v\n", len(strSlice), cap(strSlice))

		newStr[0] = "china"
		fmt.Printf("newStr: %v\n", newStr)
		fmt.Printf("oldStr: %v\n", oldStr)

		newStr = append(newStr, "brazil")
		fmt.Printf("newStr: %v\n", newStr)
		fmt.Printf("oldStr: %v\n", oldStr)

		oldStr = append(oldStr, "us")
		newStr = append(newStr, "uk")
		fmt.Printf("newStr: %v\n", newStr)
		fmt.Printf("oldStr: %v\n", oldStr)
	case "append to existing":
		var slice1 = []string{"india", "japan", "canada"}
		var slice2 = []string{"australia", "russia"}
		slice2 = append(slice2, slice1...)
		fmt.Printf("slice1: %v\n", slice1)
		fmt.Printf("slice2: %v\n", slice2)
	default:
		panic(fmt.Sprintf("uh oh", input))
	}
}
