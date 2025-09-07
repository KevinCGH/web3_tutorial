package main

func PlusOne(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < length; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	digits = make([]int, length+1)
	digits[0] = 1
	return digits
}
