package main

/**
 * 只出现一次的数字
 * https://leetcode.cn/problems/single-number/
 **/
func SingleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
