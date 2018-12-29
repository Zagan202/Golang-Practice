package main

import "fmt"

func mapExamples() {
	fmt.Println("Type: empty \nA Empty Map declaration. Map empty is created having string as key-type and int as value-type")
	fmt.Println("Type: init \nMap initialization" +
		"\nThe literal mapped values are specified using a colon-separated pair of key and value.")
	fmt.Println("Type: make \nThe make function takes as argument the type of the map and it returns an initialized map.")
	fmt.Println("Type: len \nThe built-in len() function returns the number of elements in a map." +
		"\nThe built-in len() function returns the number of elements in a map.")
	fmt.Println("Type: delete \nThe built-in delete function deletes an element from a given map associated with the provided key.")
	fmt.Println("Type: add and edit\n2 elements added and 1 edited in employee map after initialization.")
	fmt.Println("Type: contains\nAn if statment that verifies if a Map contains a key and returns a value.")
	fmt.Println("Type: iterate\nThe for…range loop statement can be used to fetch the index and element of a map.")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "empty":
		var empty = map[string]int{}
		fmt.Println(empty)
	case "init":
		var employee = map[string]int{"Mark": 10, "Sandy": 20}
		fmt.Println(employee)
	case "make":
		var employee = make(map[string]int)
		employee["Mark"] = 10
		employee["Sandy"] = 20
		fmt.Println(employee)

		employeeList := make(map[string]int)
		employeeList["Mark"] = 10
		employeeList["Sandy"] = 20
		fmt.Println(employeeList)
	case "len":
		employee := make(map[string]int)
		employee["Mark"] = 10
		employee["Sandy"] = 20

		employeeList := make(map[string]int)
		fmt.Println(len(employee))
		fmt.Println(len(employeeList))
	case "delete":
		var employee = make(map[string]int)
		employee["Mark"] = 10
		employee["Sandy"] = 20
		employee["Rocky"] = 30
		employee["Josef"] = 40

		fmt.Println(employee)

		delete(employee, "Mark")
		fmt.Println(employee)
	case "add and edit":
		var employee = map[string]int{"Mark": 10, "Sandy": 20}
		fmt.Println(employee) // Initial Map
		fmt.Println("Adding key val Rocky = 30")
		employee["Rocky"] = 30 // Add element
		employee["Josef"] = 40
		fmt.Println("editing key Mark")
		employee["Mark"] = 50 // Edit element
		fmt.Println(employee)
	case "contains":
		var employee = map[string]int{"Lawrence": 10, "Alex": 20}
		if val, ok := employee["Lawrence"]; ok {
			fmt.Println("employee[Lawrence] = ", val)
		}
		if val, ok := employee["Spiderman"]; ok {
			fmt.Println("employee[Spiderman] = ", val)
		} else {
			fmt.Println("GET ME SPIDERMAN!")
		}
	case "iterate":
		var employee = map[string]int{"Mark": 10, "Sandy": 20,
			"Rocky": 30, "Rajiv": 40, "Kate": 50}
		for key, element := range employee {
			fmt.Println("Key:", key, "=>", "Element:", element)
		}
	}

}
