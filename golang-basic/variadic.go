package main

import (
	"fmt"
	"reflect"
)

/*
Pass different types of arguments in variadic function
In the following example, the function signature accepts an
arbitrary number of arguments of type any type
*/
func whatType(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}

/*Normal function parameter with variadic function parameter*/
func calcVar(str string, y ...int) int {

	area := 1

	for _, val := range y {
		if str == "Rectangle" {
			area *= val
		} else if str == "Square" {
			area = val * val
		}
	}
	return area
}

func multiString(s ...string) {
	fmt.Println(s)
}
