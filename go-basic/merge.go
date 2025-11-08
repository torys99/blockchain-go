/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/
package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{}
	current := intervals[0]
	for _, interval := range intervals[1:] {
		if interval[0] <= current[1] {
			if interval[1] > current[1] {
				current[1] = interval[1]
			}
		} else {
			merged = append(merged, current)
			current = interval
		}
	}
	merged = append(merged, current)
	return merged
}
func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merged := merge(intervals)
	fmt.Println("合并后的区间: ", merged) // Output: [[1 6] [8 10] [15 18]]
	intervals = [][]int{{1, 4}, {4, 5}}
	merged = merge(intervals)
	fmt.Println("合并后的区间: ", merged) // Output: [[1 5]]
}
