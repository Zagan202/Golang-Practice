/*Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
package main

func twoSum(nums []int, target int) []int {
	var sol []int
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j] == target-nums[i] {
				sol = append(sol, i, j)
				return sol
			}
		}

	}
	return sol
}

func hashTwoSum(nums []int, target int) []int {
	var sol = make([]int, 2)
	var numMap = make(map[int]int)
	/*
		numMap
			Key = the value in the array nums
			Val = the index of the value in array nums
	*/
	for i := range nums {
		complement := target - nums[i]

		if val, ok := numMap[complement]; ok {
			sol[0] = val
			sol[1] = i
		} else {
			numMap[nums[i]] = i
		}
	}
	return sol
}
