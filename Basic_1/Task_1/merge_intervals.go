package Task_1

import "slices"

func MergeIntervals(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	for _, interval := range intervals {
		m := len(ans)
		if m > 0 && interval[0] <= ans[m-1][1] { // 新区间左端点 <= 栈顶区间右端点
			ans[m-1][1] = max(ans[m-1][1], interval[1])
		} else {
			ans = append(ans, interval)
		}
	}

	return ans
}
