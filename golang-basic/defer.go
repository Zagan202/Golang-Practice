package main

import "fmt"

func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}
func printing() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

//This function will call the other ones
func Defer() {
	//defer second()
	//first()
	defer printing()
}
