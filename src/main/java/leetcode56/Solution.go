package main

import (
	"fmt"
	"sort"
)

/*

给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
https://leetcode-cn.com/problems/merge-intervals/
*/
func main() {
	//[[1,3],[2,6],[8,10],[15,18]]
	//[[2,3],[4,5],[6,7],[8,9],[1,10]]
	//[[2,3],[5,5],[2,2],[3,4],[3,4]]
	//[[5,7],[2,4],[4,6],[1,1],[2,2]]
	//[[5,6],[4,4],[3,3],[2,2],[5,5]]
	fmt.Println(merge([][]int{{1, 3}, {2, 7}, {2, 6}, {15, 18}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {0, 4}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
	fmt.Println(merge([][]int{{2, 3}, {5, 5}, {2, 2}, {3, 4}, {3, 4}}))
	fmt.Println(merge([][]int{{5, 7}, {2, 4}, {4, 6}, {1, 1}, {2, 2}}))
	fmt.Println(merge([][]int{{5, 6}, {4, 4}, {3, 3}, {2, 2}, {5, 5}}))
}
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{}
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		m := merged[len(merged)-1]
		c := intervals[i]

		if c[0] > m[1] {
			merged = append(merged, c)
			continue
		}

		if c[1] > m[1] {
			m[1] = c[1]
		}
	}

	return merged
}
