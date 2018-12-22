/*
This is simply going over the way with variables
as a statically typed language
*/

package main

//https://talks.golang.org/2014/organizeio.slide#3
/*
All Go source is part of a package.
Every file begins with a package statement.
Programs start in package main.

For very small programs, main is the only package you need to write.
*/

import "fmt"

//https://golang.org/pkg/fmt/
//fmt package implements I/O with functions analogous to C's printf and scanf.

var (
	a int
	b int
)

//Globally declaring two ints "a" & "b"

func datatype() {
	var x int = 1
	var y int
	//integer data types
	fmt.Println(x) //Should print 1
	fmt.Println(y) //Should print 0

	var t, e, s, ts = 1.2, 2.3, 3.4, 4.5
	//Declaring multiple float32 vars
	fmt.Println(t, e, s, ts)

	location := "Los Angeles" //String declaration
	/*
		https://tour.golang.org/basics/10
		Short variable declarations
		:= short assignment statement can be used in
		place of a var declaration with implicit type.
	*/
	Country := "USA! USA! USA!"
	fmt.Println(location)
	fmt.Println(Country)

	alice, bob, jeff := "String", 20, 3.4 //Multiple type declarations using short assignment
	fmt.Println(alice, bob, jeff)

	a, b = 12, 43
	fmt.Println(a, b)
}
