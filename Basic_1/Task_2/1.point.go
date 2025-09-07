package Task_2

func incrementByTen(num *int) {
	*num += 10
}
func multiplySliceItemByTwo(nums *[]int) {
	for i := 0; i < len(*nums); i++ {
		(*nums)[i] *= 2
	}
}
