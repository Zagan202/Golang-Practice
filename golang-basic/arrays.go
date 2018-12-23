package main

import "fmt"

func arrayStuff() {
	var input string
	fmt.Scanln(&input)

	switch input {
	case "basic":
		var intArray [1]int
		intArray[0] = 10
		fmt.Println(intArray[0])
	case "literal 1":
		var intArray = [5]int{10, 20, 30}
		for i := 0; i < len(intArray); i++ {
			fmt.Println(intArray[i])
		}
	case "literal 2":
		//Declare an array with some or all values to it at the time of declaration
		var intArray = [5]int{0: 10, 2: 30, 4: 50}
		for i := 0; i < len(intArray); i++ {
			fmt.Println(intArray[i])
		}
	case "literal 3":
		//Declaring an array with ellipses ...
		intArray := [...]int{10, 20, 30, 40, 50}
		for i := 0; i < len(intArray); i++ {
			fmt.Println(intArray[i])
		}
	case "explore type":
		var x [5]int // Array Declaration

		x[0] = 10 // Assign the values to specific Index
		x[4] = 20 // Assign Value to array index in any Order
		x[1] = 30
		x[3] = 40
		x[2] = 50
		fmt.Println("Values of Array X: ", x)

		// Array Declartion and Intialization to specific Index
		y := [5]int{0: 100, 1: 200, 3: 500}
		fmt.Println("Values of Array Y: ", y)

		// Array Declartion and Intialization
		Country := [5]string{"US", "UK", "Australia", "Russia", "Brazil"}
		fmt.Println("Values of Array Country: ", Country)

		// Array Declartion without length and Intialization
		Transport := [...]string{"Train", "Bus", "Plane", "Car", "Bike"}
		fmt.Println("Values of Array Transport: ", Transport)
	default:
		panic(fmt.Sprintf("uh oh", input))
	}
}
