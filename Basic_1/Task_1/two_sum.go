package main

func TwoSum(nums []int, target int) []int {
	hashtable := make(map[int]int)
	for i, num := range nums {
		if p, ok := hashtable[target-num]; ok {
			return []int{p, i}
		}
		hashtable[num] = i
	}
	return []int{}
}
