package main

import (
	"math"
)

func isPalindrome(x int) bool {
	tmp := x
	var result int
	for tmp != 0 {
		if (result*10)/10 != result {
			return false
		}
		result = result*10 + tmp%10
		tmp = tmp / 10
	}
	//fmt.Println("x =", x, " result=", result)
	if result > math.MaxInt32 || result < math.MinInt32 || result != x || x < 0 {
		return false
	}
	return true
}
