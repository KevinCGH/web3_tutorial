package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTask1(t *testing.T) {

	Convey("Task1", t, func() {
		Convey("SingleNumber", func() {
			testcases := []struct {
				input    []int
				expected int
			}{
				{input: []int{2, 2, 1}, expected: 1},
				{input: []int{4, 1, 2, 1, 2}, expected: 4},
				{input: []int{1}, expected: 1},
			}
			for _, tc := range testcases {
				actual := SingleNumber(tc.input)
				So(actual, ShouldEqual, tc.expected)
			}
		})

		Convey("PalindromeNumber", func() {
			testcases := []struct {
				input    int
				expected bool
			}{
				{input: 121, expected: true},
				{input: -121, expected: false},
				{input: 10, expected: false},
			}
			for _, tc := range testcases {
				actual := IsPalindromeNumber(tc.input)
				So(actual, ShouldEqual, tc.expected)
			}
		})

		Convey("ValidParentheses", func() {
			testcases := []struct {
				input    string
				expected bool
			}{
				{input: "()", expected: true},
				{input: "()[]{}", expected: true},
				{input: "(]", expected: false},
				{input: "([])", expected: true},
				{input: "([)]", expected: false},
			}
			for _, tc := range testcases {
				actual := ValidParentheses(tc.input)
				So(actual, ShouldEqual, tc.expected)
			}
		})

		Convey("LongestCommonPrefix", func() {
			testcases := []struct {
				input    []string
				expected string
			}{
				{input: []string{"flower", "flow", "flight"}, expected: "fl"},
				{input: []string{"dog", "racecar", "car"}, expected: ""},
			}
			for _, tc := range testcases {
				actual := LongestCommonPrefix(tc.input)
				So(actual, ShouldEqual, tc.expected)
			}
		})

		Convey("PlusOne", func() {
			testcases := []struct {
				input    []int
				expected []int
			}{
				{input: []int{1, 2, 3}, expected: []int{1, 2, 4}},
				{input: []int{4, 3, 2, 1}, expected: []int{4, 3, 2, 2}},
				{input: []int{9}, expected: []int{1, 0}},
				{input: []int{9, 9}, expected: []int{1, 0, 0}},
				{input: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6},
					expected: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7}},
			}
			for _, tc := range testcases {
				actual := PlusOne(tc.input)
				So(actual, ShouldEqual, tc.expected)
			}
		})

		Convey("RemoveDuplicatesFromSortedArray", func() {
			testcases := []struct {
				input    []int
				expected int
			}{
				{input: []int{1, 1, 2}, expected: 2},
				{input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, expected: 5},
			}
			for _, tc := range testcases {
				actual := RemoveDuplicatesFromSortedArray(tc.input)
				So(actual, ShouldEqual, tc.expected)
			}
		})

		Convey("MergeIntervals", func() {
			testcases := []struct {
				input    [][]int
				expected [][]int
			}{
				{input: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, expected: [][]int{{1, 6}, {8, 10}, {15, 18}}},
				{input: [][]int{{1, 4}, {4, 5}}, expected: [][]int{{1, 5}}},
				{input: [][]int{{4, 7}, {1, 4}}, expected: [][]int{{1, 7}}},
			}
			for _, tc := range testcases {
				actual := MergeIntervals(tc.input)
				So(actual, ShouldEqual, tc.expected)
			}
		})

		Convey("TwoSum", func() {
			testcases := []struct {
				nums     []int
				target   int
				expected []int
			}{
				{nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
				{nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
				{nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
			}
			for _, tc := range testcases {
				actual := TwoSum(tc.nums, tc.target)
				So(actual, ShouldEqual, tc.expected)
			}
		})

	})

}
