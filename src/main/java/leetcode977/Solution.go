package main

import "fmt"

/*
给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。



示例 1：

输入：[-4,-1,0,3,10]
输出：[0,1,9,16,100]
示例 2：

输入：[-7,-3,2,3,11]
输出：[4,9,9,49,121]


提示：

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A 已按非递减顺序排序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/squares-of-a-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
func sortedSquares(A []int) []int {
	ans := make([]int, len(A))
	l := 0
	r := len(A) - 1
	i := len(ans) - 1
	for i >= 0 {
		l2 := A[l] * A[l]
		r2 := A[r] * A[r]
		if l2 > r2 {
			ans[i] = l2
			l++
		} else {
			ans[i] = r2
			r--
		}
		i--
	}
	return ans
}
