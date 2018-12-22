package main

import "fmt"

func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}
func Defer() {
	defer second()
	first()
}
