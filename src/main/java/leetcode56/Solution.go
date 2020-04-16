package main

import "fmt"

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
	//fmt.Println(merge([][]int{{1, 3},{2, 7}, {2, 6},{15, 18}, {8, 10}, {15, 18}}))
	//fmt.Println(merge([][]int{{1, 4},{0, 4}}))
	//fmt.Println(merge([][]int{{2, 3},{4, 5}, {6, 7},{8, 9}, {1, 10}}))
	fmt.Println(merge([][]int{{2, 3}, {5, 5}, {2, 2}, {3, 4}, {3, 4}}))
}
func merge2(intervals [][]int) [][]int {
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
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	return sort(intervals)
}
func sort(r [][]int) [][]int {
	length := len(r)
	if length <= 1 {
		return r
	}
	num := length / 2
	left := sort(r[:num])
	right := sort(r[num:])
	return mm(left, right)
}

func mm(intervals1 [][]int, intervals2 [][]int) [][]int {
	res := make([][]int, 0, len(intervals1)+len(intervals2))
	min := min(len(intervals1), len(intervals2))
	for i := 0; i < min; i++ {
		m1, m2 := m(intervals1[i], intervals2[i])
		if m2 == nil {
			res = append(res, m1)
		} else {
			res = append(res, m1, m2)
		}
	}

	if len(intervals1) > min {
		for i := min; i < len(intervals1); i++ {
			m1, m2 := m(intervals1[i], res[len(res)-1])
			res = res[:len(res)-1]
			if m2 == nil {
				res = append(res, m1)
			} else {
				//if m1[0] == res[len(res)-1][0]&&m1[1] == res[len(res)-1][1] {
				//	res = append(res, m2)
				//}
				//if m2[0] == res[len(res)-1][0]&&m2[1] == res[len(res)-1][1] {
				//	res = append(res, m1)
				//}
				res = append(res, m1, m2)
			}
		}
		//res = append(res, intervals1[min:]...)
	}
	if len(intervals2) > min {
		for i := min; i < len(intervals2); i++ {
			m1, m2 := m(intervals2[i], res[len(res)-1])
			res = res[:len(res)-1]
			if m2 == nil {
				res = append(res, m1)
			} else {
				res = append(res, m1, m2)
				//if m1[0] == res[len(res)-1][0]&&m1[1] == res[len(res)-1][1] {
				//	res = append(res, m2)
				//}
				//if m2[0] == res[len(res)-1][0]&&m2[1] == res[len(res)-1][1] {
				//	res = append(res, m1)
				//}
			}
		}
		//res = append(res, intervals2[min:]...)
	}
	return res
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func m(m1, m2 []int) ([]int, []int) {
	l1 := m1[0]
	r1 := m1[1]
	l2 := m2[0]
	r2 := m2[1]
	if (r1 >= l2 && l1 <= r2) || r2 >= l1 && l2 <= r1 {
		//if r1>r2 {
		//	return []int{m1[0], m1[len(m1)-1]}, nil
		//}
		//return []int{m1[0], m2[len(m2)-1]}, nil
		return []int{min(l1, l2), max(r1, r2)}, nil
	}
	//if r2 >= l1&&l2<=r1 {
	//	if m1[len(m1)-1]>m2[len(m2)-1] {
	//		return []int{m2[0], m1[len(m1)-1]}, nil
	//	}
	//	return []int{m2[0], m2[len(m2)-1]}, nil
	//}
	if m1[0] > m2[0] {
		return m2, m1
	}
	return m1, m2
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
