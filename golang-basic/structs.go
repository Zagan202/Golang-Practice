package main

import (
	"encoding/json"
	"fmt"
)

type rectangle struct {
	length int
	width  int
	color  string
	geo    geometry
}

type geometry struct {
	area      int
	perimeter int
}

type Employee struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

//Adding methods to struct types
type Salary struct {
	Basic, HRA, TA float64
}

type OtherEmployee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func (e OtherEmployee) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("===================")
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}
	return "----------------------"
}

func methodStructEX() {
	e := OtherEmployee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}
	fmt.Println(e.EmpInfo())
}

func jsonStruct() {
	jsonstring := `
    {
        "firstname": "Rocky",
        "lastname": "Sting",
        "city": "London"
    }`

	emp1 := new(Employee)
	json.Unmarshal([]byte(jsonstring), emp1)
	fmt.Println(emp1)

	emp2 := new(Employee)
	emp2.FirstName = "Ramesh"
	emp2.LastName = "Soni"
	emp2.City = "Mumbai"
	jsonStr, _ := json.Marshal(emp2)
	fmt.Printf("%s\n", jsonStr)
}

func dot() {
	fmt.Println("A struct uses a selector expression (or dot notation) to access the values stored in fields." +
		"\nselectors dot can be used as a rect.geometry.perimeter chain to fetch field values that are nested inside a struct.")
	var rect rectangle
	rect.length = 10
	rect.width = 20
	rect.color = "Green"

	rect.geo.area = rect.length * rect.width
	rect.geo.perimeter = 2 * (rect.length + rect.width)

	fmt.Println(rect)
	fmt.Println("Area:\t", rect.geo.area)
	fmt.Println("Perimeter:", rect.geo.perimeter)
}

func structAsgn() {
	fmt.Println(" illustrates struct instantiating using var and :=, :" +
		"\nYou can skip value when you are specified fields name.")
	var geo = geometry{100, 200}
	var rect1 = rectangle{10, 20, "Green", geo}
	fmt.Println(rect1)

	var rect2 = rectangle{length: 10, color: "Green"} // width value skipped
	fmt.Println(rect2)

	rect3 := rectangle{10, 20, "Green", geo}
	fmt.Println(rect3)

	rect4 := rectangle{length: 10, width: 20, color: "Green"}
	fmt.Println(rect4)

	rect5 := rectangle{width: 20, color: "Green"} // length value skipped
	fmt.Println(rect5)
}

func newInit() {
	fmt.Println("Struct instantiating using new keyword")
	rect1 := new(rectangle)
	rect1.length = 10
	rect1.width = 20
	rect1.color = "Green"
	fmt.Println(rect1)

	var rect2 = new(rectangle)
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2) // breadth skipped
}

func amperSand() {
	fmt.Println("Struct instantiating using &" +
		"when you instantiating like var rect1 = &rectangle{10,20,Green, geometry{200, 100}}" +
		"\nyou have to pass all field values and as per there data type otherwise compiler will give error like too few values in struct initializer" +
		"\nor cannot use XXXXXX (type string) as type int in field value.")
	var rect1 = &rectangle{10, 20, "Green", geometry{200, 100}} // Can't skip any value
	fmt.Println(rect1)

	var rect2 = &rectangle{}
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2) // breadth skipped

	var rect3 = &rectangle{}
	(*rect3).width = 10
	(*rect3).color = "Blue"
	fmt.Println(rect3) // length skipped
}
