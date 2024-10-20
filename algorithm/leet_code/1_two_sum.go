package leetcode

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

var nums = []int{2, 7, 11, 15}
var target = 9

func twoSum(nums []int, target int) []int {
	completed := make(map[int]int)
	for i, num := range nums {
		complete := target - num

		if index, ok := completed[complete]; ok {
			return []int{index, i}
		}
		completed[num] = i
	}
	return nil
}
