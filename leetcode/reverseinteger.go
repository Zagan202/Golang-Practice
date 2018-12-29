package main

import (
	"math"
)

func reverse(x int) int {
	var result int
	for x != 0 {
		if (result*10)/10 != result {
			return 0
		}
		result = result*10 + x%10
		x = x / 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
